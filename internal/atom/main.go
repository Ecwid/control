package atom

// Atom JS functions
const (
	ScrollIntoView   = `async function(){return await async function(a){const b=await new Promise(b=>{const c=new IntersectionObserver(a=>{b(a[0].intersectionRatio),c.disconnect()});c.observe(a)});1!==b&&a.scrollIntoView({block:"center",inline:"center",behavior:"instant"})}(this)}`
	ClearInput       = `function(){("INPUT"===this.nodeName||"TEXTAREA"===this.nodeName)?this.value="":this.innerText=""}`
	GetInnerText     = `function(){return this.value||this.innerText}`
	DispatchEvents   = `function(l){for(const e of l)this.dispatchEvent(new Event(e,{'bubbles':!0}))}`
	Select           = `function(a){const b=Array.from(this.options);this.value=void 0;for(const c of b)if(c.selected=a.includes(c.value),c.selected&&!this.multiple)break}`
	GetSelected      = `function(){return Array.from(this.options).filter(a=>a.selected).map(a=>a.value)}`
	GetSelectedText  = `function(){return Array.from(this.options).filter(a=>a.selected).map(a=>a.innerText)}`
	SelectHasOptions = `function(c){const a=Array.from(this.options);return c.length==a.filter(a=>c.includes(a.value)).length}`
	CheckBox         = `function(c){this.checked = c}`
	IsChecked        = `function(){return this.checked}`
	GetComputedStyle = `function(s){return getComputedStyle(this)[s]}`
	SetAttr          = `function(a,v){this.setAttribute(a,v)}`
	GetAttr          = `function(a){return this.getAttribute(a)}`
	IsVisible        = `function(){const b=this.getBoundingClientRect(),c=window.getComputedStyle(this);return c&&"hidden"!==c.visibility&&!c.disabled&&!!(b.top||b.bottom||b.width||b.height)}`
	Query            = `function(s){return this.querySelector(s)}`
	QueryAll         = `function(s){return this.querySelectorAll(s)}`
	IsClickHit       = `function(){return this._cc}`
	PreventMissClick = `function(){this._cc=!1,tt=this,z=function(b){for(var c=b;c;c=c.parentNode)if(c==tt)return!0;return!1},i=function(b){if (z(b.target)) {tt._cc=!0;} else {b.stopPropagation();b.preventDefault()}},document.addEventListener("click",i,{capture:!0,once:!0})}`
	MutationObserver = `function(b,d,c){return new Promise(e=>{const f=new MutationObserver(b=>{for(var c of b){e(c.type),f.disconnect();break}});f.observe(this,{attributes:b,childList:d,subtree:c})})}`
	LayoutMetrics    = `({
		width: Math.min(
		   window.innerWidth,
		   document.body.scrollWidth,
		   document.documentElement.scrollWidth) | 0,
		height: Math.max(
		   document.body.scrollHeight,
		   document.body.offsetHeight,
		   document.documentElement.clientHeight,
		   document.documentElement.scrollHeight,
		   document.documentElement.offsetHeight) | 0,
		deviceScaleFactor: window.devicePixelRatio || 1,
		mobile: typeof window.orientation !== 'undefined'
	 })`
)
