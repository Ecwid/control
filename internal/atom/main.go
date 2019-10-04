package atom

// Atom JS functions
const (
	AddEventFired    = `function(e){this._g_e=!1,this.addEventListener(e,function(){this._g_e=!0},{capture:!0,once:!0,passive:!0})}`
	IsEventFired     = `function(){return this._g_e}`
	ScrollIntoView   = `async function(){return await async function(a){const b=await new Promise(b=>{const c=new IntersectionObserver(a=>{b(a[0].intersectionRatio),c.disconnect()});c.observe(a)});1!==b&&a.scrollIntoView({block:"center",inline:"center",behavior:"instant"})}(this)}`
	IsClickableAt    = `function(a,b){for(var c=this,d=this;d.parentNode;)d=d.parentNode;return function(){var e=d.elementFromPoint(a,b);if(!e)return!1;if(e==c)return!0;for(var f=e.parentNode;f;){if(f==c)return!0;f=f.parentNode}return!1}()}`
	ClearInput       = `function(){"INPUT"===this.nodeName?this.value="":this.innerText=""}`
	GetInnerText     = `function(){return this.value||this.innerText}`
	DispatchEvents   = `function(l){for(const e of l)this.dispatchEvent(new Event(e,{'bubbles':!0}))}`
	Select           = `function(a){const b=Array.from(this.options);this.value=void 0;for(const c of b)if(c.selected=a.includes(c.value),c.selected&&!this.multiple)break;this.dispatchEvent(new Event("input",{bubbles:!0})),this.dispatchEvent(new Event("change",{bubbles:!0}))}`
	GetSelected      = `function(){return Array.from(this.options).filter(a=>a.selected).map(a=>a.value)}`
	GetSelectedText  = `function(){return Array.from(this.options).filter(a=>a.selected).map(a=>a.innerText)}`
	SelectHasOptions = `function(c){const a=Array.from(this.options);return c.length==a.filter(a=>c.includes(a.value)).length}`
	CheckBox         = `function(c){this.checked = c}`
	IsChecked        = `function(){return this.checked}`
	GetComputedStyle = `function(s){return getComputedStyle(this)[s]}`
	SetAttr          = `function(a,v){this.setAttribute(a,v)}`
	GetAttr          = `function(a){return this.getAttribute(a)}`
	IsFocusable      = `function(){const b=this.getBoundingClientRect(),c=window.getComputedStyle(this);return c&&"hidden"!==c.visibility&&!c.disabled&&!!(b.top||b.bottom||b.width||b.height)}`
	Query            = `function(s){return this.querySelector(s)}`
	QueryAll         = `function(s){return this.querySelectorAll(s)}`
)
