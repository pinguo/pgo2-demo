package pgo2

import (
	"github.com/pinguo/pgo2/iface"
)

// Object base class of context based object
type Object struct {
	context iface.IContext
}

// GetContext get context of this object
func (o *Object) Context() iface.IContext {
	return o.context
}

// SetContext set context of this object
func (o *Object) SetContext(ctx iface.IContext) {
	o.context = ctx
}

// GetObject create new object
func (o *Object) GetObj(obj iface.IObject) iface.IObject {
	obj.SetContext(o.Context())
	return obj
}

// GetObjPool Get Object from pool
func (o *Object) GetObjPool(funcName iface.IObjPoolFunc, params ...interface{}) iface.IObject {
	return funcName(o.Context(), params...)
}

// GetObject Get single object
func (o *Object) GetObjSingle(name string, funcName iface.IObjSingleFunc, params ...interface{}) iface.IObject {
	// obj := funcName(params...)
	obj := App().GetObjSingle(name, funcName, params...)
	obj.SetContext(o.Context())
	return obj
}

// GetObjPoolCtx Get Object from pool and new Context
func (o *Object) GetObjPoolCtx(ctx iface.IContext, funcName iface.IObjPoolFunc, params ...interface{}) iface.IObject {
	return funcName(ctx, params...)
}

// GetObject create new object  and new Context
func (o *Object) GetObjCtx(ctx iface.IContext, obj iface.IObject) iface.IObject {
	obj.SetContext(ctx)
	return obj
}

// GetObject Get single object and new Context
func (o *Object) GetObjSingleCtx(ctx iface.IContext, name string, funcName iface.IObjSingleFunc, params ...interface{}) iface.IObject {
	// obj := funcName(params...)
	obj := App().GetObjSingle(name, funcName, params...)
	obj.SetContext(ctx)
	return obj
}
