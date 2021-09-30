package control

const (
	WebEventKeypress = "keypress"
	WebEventKeyup    = "keyup"
	WebEventChange   = "change"
	WebEventInput    = "input"
	WebEventClick    = "click"
)

// Atom JS functions
const (
	functionClearText            = `function(){("INPUT"===this.nodeName||"TEXTAREA"===this.nodeName)?this.value="":this.innerText=""}`
	functionGetText              = `function(){switch(this.tagName){case"INPUT":case"TEXTAREA":return this.value;case"SELECT":return Array.from(this.selectedOptions).map(b=>b.innerText).join();default:return this.innerText||this.textContent;}}`
	functionDispatchEvents       = `function(l){for(const e of l)this.dispatchEvent(new Event(e,{'bubbles':!0}))}`
	functionPreventMissClick     = `function(){tt=this,z=function(b){for(var c=b;c;c=c.parentNode)if(c==tt)return!0;return!1},i=function(b){if (z(b.target)) {_control_click("1")} else {b.stopPropagation();b.preventDefault();_control_click(b.target.outerHTML||"0")}},document.addEventListener("click",i,{capture:!0,once:!0})}`
	functionSetAttr              = `function(a,v){this.setAttribute(a,v)}`
	functionGetAttr              = `function(a){return this.getAttribute(a)}`
	functionCheckbox             = `function(v){this.checked=v}`
	functionIsChecked            = `function(){return this.checked}`
	functionGetComputedStyle     = `function(s){return getComputedStyle(this)[s]}`
	functionSelect               = `function(a){const b=Array.from(this.options);this.value=void 0;for(const c of b)if(c.selected=a.includes(c.value),c.selected&&!this.multiple)break}`
	functionGetSelectedValues    = `function(){return Array.from(this.options).filter(a=>a.selected).map(a=>a.value)}`
	functionGetSelectedInnerText = `function(){return Array.from(this.options).filter(a=>a.selected).map(a=>a.innerText)}`
	functionDOMIdle              = `var d=function(e,t,n){var u,r=null;return function(){var i=this,o=arguments,s=n&&!r;return clearTimeout(r),r=setTimeout(function(){r=null,n||(u=e.apply(i,o))},t),s&&(u=e.apply(i,o)),u}};new Promise((e,t)=>{var n=d(function(){e()},%d);new MutationObserver(n).observe(document,{attributes:!0,childList:!0,subtree:!0}),n(),setTimeout(()=>t("timeout"),%d)});`
)
