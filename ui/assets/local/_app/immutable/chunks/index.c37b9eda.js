function N(){}const X=t=>t;function Et(t,e){for(const n in e)t[n]=e[n];return t}function Nt(t){return!!t&&(typeof t=="object"||typeof t=="function")&&typeof t.then=="function"}function ft(t){return t()}function at(){return Object.create(null)}function C(t){t.forEach(ft)}function T(t){return typeof t=="function"}function Xt(t,e){return t!=t?e==e:t!==e||t&&typeof t=="object"||typeof t=="function"}let L;function Yt(t,e){return L||(L=document.createElement("a")),L.href=e,t===L.href}function At(t){return Object.keys(t).length===0}function _t(t,...e){if(t==null)return N;const n=t.subscribe(...e);return n.unsubscribe?()=>n.unsubscribe():n}function Zt(t){let e;return _t(t,n=>e=n)(),e}function te(t,e,n){t.$$.on_destroy.push(_t(e,n))}function ee(t,e,n,i){if(t){const s=dt(t,e,n,i);return t[0](s)}}function dt(t,e,n,i){return t[1]&&i?Et(n.ctx.slice(),t[1](i(e))):n.ctx}function ne(t,e,n,i){if(t[2]&&i){const s=t[2](i(n));if(e.dirty===void 0)return s;if(typeof s=="object"){const c=[],r=Math.max(e.dirty.length,s.length);for(let o=0;o<r;o+=1)c[o]=e.dirty[o]|s[o];return c}return e.dirty|s}return e.dirty}function ie(t,e,n,i,s,c){if(s){const r=dt(e,n,i,c);t.p(r,s)}}function se(t){if(t.ctx.length>32){const e=[],n=t.ctx.length/32;for(let i=0;i<n;i++)e[i]=-1;return e}return-1}function re(t){const e={};for(const n in t)n[0]!=="$"&&(e[n]=t[n]);return e}function ce(t,e){const n={};e=new Set(e);for(const i in t)!e.has(i)&&i[0]!=="$"&&(n[i]=t[i]);return n}function oe(t){const e={};for(const n in t)e[n]=!0;return e}function ae(t){return t??""}function le(t,e,n){return t.set(n),e}const St=(t,e)=>Object.prototype.hasOwnProperty.call(t,e);function ue(t){return t&&T(t.destroy)?t.destroy:N}const ht=typeof window<"u";let Y=ht?()=>window.performance.now():()=>Date.now(),Z=ht?t=>requestAnimationFrame(t):N;const M=new Set;function mt(t){M.forEach(e=>{e.c(t)||(M.delete(e),e.f())}),M.size!==0&&Z(mt)}function tt(t){let e;return M.size===0&&Z(mt),{promise:new Promise(n=>{M.add(e={c:t,f:n})}),abort(){M.delete(e)}}}let K=!1;function Ct(){K=!0}function jt(){K=!1}function Dt(t,e,n,i){for(;t<e;){const s=t+(e-t>>1);n(s)<=i?t=s+1:e=s}return t}function Mt(t){if(t.hydrate_init)return;t.hydrate_init=!0;let e=t.childNodes;if(t.nodeName==="HEAD"){const a=[];for(let l=0;l<e.length;l++){const _=e[l];_.claim_order!==void 0&&a.push(_)}e=a}const n=new Int32Array(e.length+1),i=new Int32Array(e.length);n[0]=-1;let s=0;for(let a=0;a<e.length;a++){const l=e[a].claim_order,_=(s>0&&e[n[s]].claim_order<=l?s+1:Dt(1,s,d=>e[n[d]].claim_order,l))-1;i[a]=n[_]+1;const f=_+1;n[f]=a,s=Math.max(f,s)}const c=[],r=[];let o=e.length-1;for(let a=n[s]+1;a!=0;a=i[a-1]){for(c.push(e[a-1]);o>=a;o--)r.push(e[o]);o--}for(;o>=0;o--)r.push(e[o]);c.reverse(),r.sort((a,l)=>a.claim_order-l.claim_order);for(let a=0,l=0;a<r.length;a++){for(;l<c.length&&r[a].claim_order>=c[l].claim_order;)l++;const _=l<c.length?c[l]:null;t.insertBefore(r[a],_)}}function pt(t,e){t.appendChild(e)}function yt(t){if(!t)return document;const e=t.getRootNode?t.getRootNode():t.ownerDocument;return e&&e.host?e:t.ownerDocument}function Ot(t){const e=nt("style");return Tt(yt(t),e),e.sheet}function Tt(t,e){return pt(t.head||t,e),e.sheet}function Pt(t,e){if(K){for(Mt(t),(t.actual_end_child===void 0||t.actual_end_child!==null&&t.actual_end_child.parentNode!==t)&&(t.actual_end_child=t.firstChild);t.actual_end_child!==null&&t.actual_end_child.claim_order===void 0;)t.actual_end_child=t.actual_end_child.nextSibling;e!==t.actual_end_child?(e.claim_order!==void 0||e.parentNode!==t)&&t.insertBefore(e,t.actual_end_child):t.actual_end_child=e.nextSibling}else(e.parentNode!==t||e.nextSibling!==null)&&t.appendChild(e)}function fe(t,e,n){K&&!n?Pt(t,e):(e.parentNode!==t||e.nextSibling!=n)&&t.insertBefore(e,n||null)}function et(t){t.parentNode&&t.parentNode.removeChild(t)}function _e(t,e){for(let n=0;n<t.length;n+=1)t[n]&&t[n].d(e)}function nt(t){return document.createElement(t)}function de(t,e){const n={};for(const i in t)St(t,i)&&e.indexOf(i)===-1&&(n[i]=t[i]);return n}function zt(t){return document.createElementNS("http://www.w3.org/2000/svg",t)}function it(t){return document.createTextNode(t)}function he(){return it(" ")}function me(){return it("")}function lt(t,e,n,i){return t.addEventListener(e,n,i),()=>t.removeEventListener(e,n,i)}function pe(t){return function(e){return e.preventDefault(),t.call(this,e)}}function ye(t){return function(e){return e.stopPropagation(),t.call(this,e)}}function gt(t,e,n){n==null?t.removeAttribute(e):t.getAttribute(e)!==n&&t.setAttribute(e,n)}function ge(t,e){const n=Object.getOwnPropertyDescriptors(t.__proto__);for(const i in e)e[i]==null?t.removeAttribute(i):i==="style"?t.style.cssText=e[i]:i==="__value"?t.value=t[i]=e[i]:n[i]&&n[i].set?t[i]=e[i]:gt(t,i,e[i])}function be(t,e){Object.keys(e).forEach(n=>{qt(t,n,e[n])})}function qt(t,e,n){e in t?t[e]=typeof t[e]=="boolean"&&n===""?!0:n:gt(t,e,n)}function we(t){return t===""?null:+t}function Rt(t){return Array.from(t.childNodes)}function Lt(t){t.claim_info===void 0&&(t.claim_info={last_index:0,total_claimed:0})}function bt(t,e,n,i,s=!1){Lt(t);const c=(()=>{for(let r=t.claim_info.last_index;r<t.length;r++){const o=t[r];if(e(o)){const a=n(o);return a===void 0?t.splice(r,1):t[r]=a,s||(t.claim_info.last_index=r),o}}for(let r=t.claim_info.last_index-1;r>=0;r--){const o=t[r];if(e(o)){const a=n(o);return a===void 0?t.splice(r,1):t[r]=a,s?a===void 0&&t.claim_info.last_index--:t.claim_info.last_index=r,o}}return i()})();return c.claim_order=t.claim_info.total_claimed,t.claim_info.total_claimed+=1,c}function wt(t,e,n,i){return bt(t,s=>s.nodeName===e,s=>{const c=[];for(let r=0;r<s.attributes.length;r++){const o=s.attributes[r];n[o.name]||c.push(o.name)}c.forEach(r=>s.removeAttribute(r))},()=>i(e))}function $e(t,e,n){return wt(t,e,n,nt)}function xe(t,e,n){return wt(t,e,n,zt)}function Wt(t,e){return bt(t,n=>n.nodeType===3,n=>{const i=""+e;if(n.data.startsWith(i)){if(n.data.length!==i.length)return n.splitText(i.length)}else n.data=i},()=>it(e),!0)}function ke(t){return Wt(t," ")}function ve(t,e){e=""+e,t.wholeText!==e&&(t.data=e)}function Ee(t,e){t.value=e??""}function Ne(t,e,n,i){n===null?t.style.removeProperty(e):t.style.setProperty(e,n,i?"important":"")}function Ae(t,e){for(let n=0;n<t.options.length;n+=1){const i=t.options[n];if(i.__value===e){i.selected=!0;return}}t.selectedIndex=-1}function Se(t){const e=t.querySelector(":checked")||t.options[0];return e&&e.__value}let W;function Bt(){if(W===void 0){W=!1;try{typeof window<"u"&&window.parent&&window.parent.document}catch{W=!0}}return W}function Ce(t,e){getComputedStyle(t).position==="static"&&(t.style.position="relative");const i=nt("iframe");i.setAttribute("style","display: block; position: absolute; top: 0; left: 0; width: 100%; height: 100%; overflow: hidden; border: 0; opacity: 0; pointer-events: none; z-index: -1;"),i.setAttribute("aria-hidden","true"),i.tabIndex=-1;const s=Bt();let c;return s?(i.src="data:text/html,<script>onresize=function(){parent.postMessage(0,'*')}<\/script>",c=lt(window,"message",r=>{r.source===i.contentWindow&&e()})):(i.src="about:blank",i.onload=()=>{c=lt(i.contentWindow,"resize",e)}),pt(t,i),()=>{(s||c&&i.contentWindow)&&c(),et(i)}}function je(t,e,n){t.classList[n?"add":"remove"](e)}function $t(t,e,{bubbles:n=!1,cancelable:i=!1}={}){const s=document.createEvent("CustomEvent");return s.initCustomEvent(t,n,i,e),s}function De(t,e){const n=[];let i=0;for(const s of e.childNodes)if(s.nodeType===8){const c=s.textContent.trim();c===`HEAD_${t}_END`?(i-=1,n.push(s)):c===`HEAD_${t}_START`&&(i+=1,n.push(s))}else i>0&&n.push(s);return n}function Me(t,e){return new t(e)}const F=new Map;let I=0;function Ht(t){let e=5381,n=t.length;for(;n--;)e=(e<<5)-e^t.charCodeAt(n);return e>>>0}function Ft(t,e){const n={stylesheet:Ot(e),rules:{}};return F.set(t,n),n}function G(t,e,n,i,s,c,r,o=0){const a=16.666/i;let l=`{
`;for(let m=0;m<=1;m+=a){const g=e+(n-e)*c(m);l+=m*100+`%{${r(g,1-g)}}
`}const _=l+`100% {${r(n,1-n)}}
}`,f=`__svelte_${Ht(_)}_${o}`,d=yt(t),{stylesheet:u,rules:h}=F.get(d)||Ft(d,t);h[f]||(h[f]=!0,u.insertRule(`@keyframes ${f} ${_}`,u.cssRules.length));const p=t.style.animation||"";return t.style.animation=`${p?`${p}, `:""}${f} ${i}ms linear ${s}ms 1 both`,I+=1,f}function J(t,e){const n=(t.style.animation||"").split(", "),i=n.filter(e?c=>c.indexOf(e)<0:c=>c.indexOf("__svelte")===-1),s=n.length-i.length;s&&(t.style.animation=i.join(", "),I-=s,I||It())}function It(){Z(()=>{I||(F.forEach(t=>{const{ownerNode:e}=t.stylesheet;e&&et(e)}),F.clear())})}let q;function E(t){q=t}function A(){if(!q)throw new Error("Function called outside component initialization");return q}function Oe(t){A().$$.on_mount.push(t)}function Te(t){A().$$.after_update.push(t)}function Pe(t){A().$$.on_destroy.push(t)}function ze(){const t=A();return(e,n,{cancelable:i=!1}={})=>{const s=t.$$.callbacks[e];if(s){const c=$t(e,n,{cancelable:i});return s.slice().forEach(r=>{r.call(t,c)}),!c.defaultPrevented}return!0}}function qe(t,e){return A().$$.context.set(t,e),e}function Re(t){return A().$$.context.get(t)}function Le(t){return A().$$.context.has(t)}function We(t,e){const n=t.$$.callbacks[e.type];n&&n.slice().forEach(i=>i.call(this,e))}const D=[],ut=[],B=[],U=[],xt=Promise.resolve();let V=!1;function kt(){V||(V=!0,xt.then(st))}function Be(){return kt(),xt}function O(t){B.push(t)}function He(t){U.push(t)}const Q=new Set;let j=0;function st(){if(j!==0)return;const t=q;do{try{for(;j<D.length;){const e=D[j];j++,E(e),Gt(e.$$)}}catch(e){throw D.length=0,j=0,e}for(E(null),D.length=0,j=0;ut.length;)ut.pop()();for(let e=0;e<B.length;e+=1){const n=B[e];Q.has(n)||(Q.add(n),n())}B.length=0}while(D.length);for(;U.length;)U.pop()();V=!1,Q.clear(),E(t)}function Gt(t){if(t.fragment!==null){t.update(),C(t.before_update);const e=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,e),t.after_update.forEach(O)}}let z;function rt(){return z||(z=Promise.resolve(),z.then(()=>{z=null})),z}function S(t,e,n){t.dispatchEvent($t(`${e?"intro":"outro"}${n}`))}const H=new Set;let k;function Jt(){k={r:0,c:[],p:k}}function Kt(){k.r||C(k.c),k=k.p}function ct(t,e){t&&t.i&&(H.delete(t),t.i(e))}function vt(t,e,n,i){if(t&&t.o){if(H.has(t))return;H.add(t),k.c.push(()=>{H.delete(t),i&&(n&&t.d(1),i())}),t.o(e)}else i&&i()}const ot={duration:0};function Fe(t,e,n){const i={direction:"in"};let s=e(t,n,i),c=!1,r,o,a=0;function l(){r&&J(t,r)}function _(){const{delay:d=0,duration:u=300,easing:h=X,tick:p=N,css:m}=s||ot;m&&(r=G(t,0,1,u,d,h,m,a++)),p(0,1);const g=Y()+d,v=g+u;o&&o.abort(),c=!0,O(()=>S(t,!0,"start")),o=tt(b=>{if(c){if(b>=v)return p(1,0),S(t,!0,"end"),l(),c=!1;if(b>=g){const w=h((b-g)/u);p(w,1-w)}}return c})}let f=!1;return{start(){f||(f=!0,J(t),T(s)?(s=s(i),rt().then(_)):_())},invalidate(){f=!1},end(){c&&(l(),c=!1)}}}function Ie(t,e,n){const i={direction:"out"};let s=e(t,n,i),c=!0,r;const o=k;o.r+=1;function a(){const{delay:l=0,duration:_=300,easing:f=X,tick:d=N,css:u}=s||ot;u&&(r=G(t,1,0,_,l,f,u));const h=Y()+l,p=h+_;O(()=>S(t,!1,"start")),tt(m=>{if(c){if(m>=p)return d(0,1),S(t,!1,"end"),--o.r||C(o.c),!1;if(m>=h){const g=f((m-h)/_);d(1-g,g)}}return c})}return T(s)?rt().then(()=>{s=s(i),a()}):a(),{end(l){l&&s.tick&&s.tick(1,0),c&&(r&&J(t,r),c=!1)}}}function Ge(t,e,n,i){const s={direction:"both"};let c=e(t,n,s),r=i?0:1,o=null,a=null,l=null;function _(){l&&J(t,l)}function f(u,h){const p=u.b-r;return h*=Math.abs(p),{a:r,b:u.b,d:p,duration:h,start:u.start,end:u.start+h,group:u.group}}function d(u){const{delay:h=0,duration:p=300,easing:m=X,tick:g=N,css:v}=c||ot,b={start:Y()+h,b:u};u||(b.group=k,k.r+=1),o||a?a=b:(v&&(_(),l=G(t,r,u,p,h,m,v)),u&&g(0,1),o=f(b,p),O(()=>S(t,u,"start")),tt(w=>{if(a&&w>a.start&&(o=f(a,p),a=null,S(t,o.b,"start"),v&&(_(),l=G(t,r,o.b,o.duration,0,m,c.css))),o){if(w>=o.end)g(r=o.b,1-r),S(t,o.b,"end"),a||(o.b?_():--o.group.r||C(o.group.c)),o=null;else if(w>=o.start){const P=w-o.start;r=o.a+o.d*m(P/o.duration),g(r,1-r)}}return!!(o||a)}))}return{run(u){T(c)?rt().then(()=>{c=c(s),d(u)}):d(u)},end(){_(),o=a=null}}}function Je(t,e){const n=e.token={};function i(s,c,r,o){if(e.token!==n)return;e.resolved=o;let a=e.ctx;r!==void 0&&(a=a.slice(),a[r]=o);const l=s&&(e.current=s)(a);let _=!1;e.block&&(e.blocks?e.blocks.forEach((f,d)=>{d!==c&&f&&(Jt(),vt(f,1,1,()=>{e.blocks[d]===f&&(e.blocks[d]=null)}),Kt())}):e.block.d(1),l.c(),ct(l,1),l.m(e.mount(),e.anchor),_=!0),e.block=l,e.blocks&&(e.blocks[c]=l),_&&st()}if(Nt(t)){const s=A();if(t.then(c=>{E(s),i(e.then,1,e.value,c),E(null)},c=>{if(E(s),i(e.catch,2,e.error,c),E(null),!e.hasCatch)throw c}),e.current!==e.pending)return i(e.pending,0),!0}else{if(e.current!==e.then)return i(e.then,1,e.value,t),!0;e.resolved=t}}function Ke(t,e,n){const i=e.slice(),{resolved:s}=t;t.current===t.then&&(i[t.value]=s),t.current===t.catch&&(i[t.error]=s),t.block.p(i,n)}const Qe=typeof window<"u"?window:typeof globalThis<"u"?globalThis:global;function Ue(t,e){t.d(1),e.delete(t.key)}function Ve(t,e){vt(t,1,1,()=>{e.delete(t.key)})}function Xe(t,e,n,i,s,c,r,o,a,l,_,f){let d=t.length,u=c.length,h=d;const p={};for(;h--;)p[t[h].key]=h;const m=[],g=new Map,v=new Map;for(h=u;h--;){const y=f(s,c,h),$=n(y);let x=r.get($);x?i&&x.p(y,e):(x=l($,y),x.c()),g.set($,m[h]=x),$ in p&&v.set($,Math.abs(h-p[$]))}const b=new Set,w=new Set;function P(y){ct(y,1),y.m(o,_),r.set(y.key,y),_=y.first,u--}for(;d&&u;){const y=m[u-1],$=t[d-1],x=y.key,R=$.key;y===$?(_=y.first,d--,u--):g.has(R)?!r.has(x)||b.has(x)?P(y):w.has(R)?d--:v.get(x)>v.get(R)?(w.add(x),P(y)):(b.add(R),d--):(a($,r),d--)}for(;d--;){const y=t[d];g.has(y.key)||a(y,r)}for(;u;)P(m[u-1]);return m}function Ye(t,e){const n={},i={},s={$$scope:1};let c=t.length;for(;c--;){const r=t[c],o=e[c];if(o){for(const a in r)a in o||(i[a]=1);for(const a in o)s[a]||(n[a]=o[a],s[a]=1);t[c]=o}else for(const a in r)s[a]=1}for(const r in i)r in n||(n[r]=void 0);return n}function Ze(t){return typeof t=="object"&&t!==null?t:{}}function tn(t,e,n){const i=t.$$.props[e];i!==void 0&&(t.$$.bound[i]=n,n(t.$$.ctx[i]))}function en(t){t&&t.c()}function nn(t,e){t&&t.l(e)}function Qt(t,e,n,i){const{fragment:s,after_update:c}=t.$$;s&&s.m(e,n),i||O(()=>{const r=t.$$.on_mount.map(ft).filter(T);t.$$.on_destroy?t.$$.on_destroy.push(...r):C(r),t.$$.on_mount=[]}),c.forEach(O)}function Ut(t,e){const n=t.$$;n.fragment!==null&&(C(n.on_destroy),n.fragment&&n.fragment.d(e),n.on_destroy=n.fragment=null,n.ctx=[])}function Vt(t,e){t.$$.dirty[0]===-1&&(D.push(t),kt(),t.$$.dirty.fill(0)),t.$$.dirty[e/31|0]|=1<<e%31}function sn(t,e,n,i,s,c,r,o=[-1]){const a=q;E(t);const l=t.$$={fragment:null,ctx:[],props:c,update:N,not_equal:s,bound:at(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(e.context||(a?a.$$.context:[])),callbacks:at(),dirty:o,skip_bound:!1,root:e.target||a.$$.root};r&&r(l.root);let _=!1;if(l.ctx=n?n(t,e.props||{},(f,d,...u)=>{const h=u.length?u[0]:d;return l.ctx&&s(l.ctx[f],l.ctx[f]=h)&&(!l.skip_bound&&l.bound[f]&&l.bound[f](h),_&&Vt(t,f)),d}):[],l.update(),_=!0,C(l.before_update),l.fragment=i?i(l.ctx):!1,e.target){if(e.hydrate){Ct();const f=Rt(e.target);l.fragment&&l.fragment.l(f),f.forEach(et)}else l.fragment&&l.fragment.c();e.intro&&ct(t.$$.fragment),Qt(t,e.target,e.anchor,e.customElement),jt(),st()}E(a)}class rn{$destroy(){Ut(this,1),this.$destroy=N}$on(e,n){if(!T(n))return N;const i=this.$$.callbacks[e]||(this.$$.callbacks[e]=[]);return i.push(n),()=>{const s=i.indexOf(n);s!==-1&&i.splice(s,1)}}$set(e){this.$$set&&!At(e)&&(this.$$.skip_bound=!0,this.$$set(e),this.$$.skip_bound=!1)}}export{O as $,ut as A,Me as B,en as C,nn as D,Qt as E,Ut as F,ee as G,ie as H,se as I,ne as J,te as K,Pt as L,ze as M,Et as N,re as O,We as P,ge as Q,je as R,rn as S,lt as T,Ye as U,ye as V,Zt as W,ae as X,tn as Y,He as Z,le as _,Xt as a,Ge as a0,Xe as a1,Ve as a2,zt as a3,xe as a4,ce as a5,Ze as a6,_e as a7,Ee as a8,ue as a9,Fe as aa,Yt as ab,qe as ac,pe as ad,Re as ae,X as af,oe as ag,De as ah,Pe as ai,be as aj,Le as ak,Ue as al,Ce as am,Ie as an,Ae as ao,Se as ap,we as aq,Je as ar,Ke as as,Qe as at,de as au,qt as av,sn as b,he as c,ke as d,me as e,fe as f,vt as g,Kt as h,T as i,ct as j,et as k,Te as l,nt as m,N as n,Oe as o,$e as p,Rt as q,C as r,_t as s,Be as t,gt as u,Ne as v,it as w,Wt as x,ve as y,Jt as z};
