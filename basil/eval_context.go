package basil

import (
	"sync"
)

// EvalStage means an evaluation stage (default, pre or post)
type EvalStage int8

// Evaluation stages
const (
	EvalStageInit  EvalStage = -1
	EvalStageMain  EvalStage = 0
	EvalStageClose EvalStage = 1
)

// EvalStages returns with the evaluation stages
var EvalStages = map[string]EvalStage{
	"init":  EvalStageInit,
	"main":  EvalStageMain,
	"close": EvalStageClose,
}

type EvalContext interface {
	WithDependencies(map[ID]BlockContainer) EvalContext
	BlockContext() BlockContext
	SetBlockContext(BlockContext)
	BlockContainer(ID) (BlockContainer, bool)
	ScheduleJob(Job)
	Subscribe(*NodeContainer, ID)
	Publish(Container)
	Logger() Logger
}

// EvalContext is the evaluation context
type evalContext struct {
	parentCtx    EvalContext
	blockCtx     BlockContext
	scheduler    Scheduler
	dependencies map[ID]BlockContainer
	pubsub       *pubsub
}

// NewEvalContext returns with a new evaluation context
func NewEvalContext(blockCtx BlockContext, scheduler Scheduler) EvalContext {
	return &evalContext{
		blockCtx:  blockCtx,
		scheduler: scheduler,
		pubsub:    newPubSub(),
	}
}

// WithDependencies returns a copy of parent with the given dependencies
func (e *evalContext) WithDependencies(dependencies map[ID]BlockContainer) EvalContext {
	return &evalContext{
		parentCtx:    e,
		blockCtx:     e.blockCtx,
		scheduler:    e.scheduler,
		pubsub:       e.pubsub,
		dependencies: dependencies,
	}
}

// BlockContext returns with the block context
func (e *evalContext) BlockContext() BlockContext {
	return e.blockCtx
}

// SetBlockContext sets the block context
func (e *evalContext) SetBlockContext(blockContext BlockContext) {
	e.blockCtx = blockContext
}

// BlockContainer returns with the given block container instance if it exists
func (e *evalContext) BlockContainer(id ID) (BlockContainer, bool) {
	if container, ok := e.dependencies[id]; ok {
		return container, true
	}

	if e.parentCtx != nil {
		return e.parentCtx.BlockContainer(id)
	}

	return nil, false
}

func (e *evalContext) ScheduleJob(job Job) {
	e.scheduler.Schedule(job)
}

func (e *evalContext) Subscribe(container *NodeContainer, id ID) {
	e.pubsub.Subscribe(container, id)
}

func (e *evalContext) Publish(c Container) {
	e.pubsub.Publish(c)
}

func (e *evalContext) Logger() Logger {
	return e.blockCtx.Logger()
}

type pubsub struct {
	subs       map[ID][]*NodeContainer
	generators map[ID][]*NodeContainer
	mu         *sync.RWMutex
}

func newPubSub() *pubsub {
	return &pubsub{
		subs:       make(map[ID][]*NodeContainer),
		generators: make(map[ID][]*NodeContainer),
		mu:         &sync.RWMutex{},
	}
}

// Subscribe will subscribe the given node container for the given dependency
func (p *pubsub) Subscribe(container *NodeContainer, id ID) {
	p.mu.Lock()
	p.subs[id] = append(p.subs[id], container)
	p.mu.Unlock()
}

// Publish will notify all node containers which are subscribed for the dependency
// The ready function will run on any containers which have all dependencies satisfied
func (p *pubsub) Publish(c Container) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, container := range p.subs[c.ID()] {
		container.SetDependency(c)
	}
}
