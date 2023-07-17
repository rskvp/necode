import{S as ke,b as Te,a as pe,G as ae,N as w,m as M,c as z,C as x,p as I,q as A,d as J,k as g,D as $,u as D,Q as ie,R as G,f as K,L as p,E as ee,T as Z,V as Ee,ad as Ce,j as E,z as ne,g as N,h as se,H as oe,I as fe,J as re,U as De,F as te,r as Ne,a5 as ce,M as ye,O as de,w as Q,x as U,y as le,A as Me,n as be}from"./index.c37b9eda.js";import{B as ve}from"./button.b208c7de.js";import{I as Ie}from"./tooltip.ddb5fe07.js";const Ae=a=>({}),ue=a=>({}),Fe=a=>({}),me=a=>({});function _e(a){let e,l;return e=new Ie({props:{"aria-label":a[2],icon:"close",class:"float-right m-4"}}),e.$on("click",a[12]),{c(){x(e.$$.fragment)},l(t){$(e.$$.fragment,t)},m(t,s){ee(e,t,s),l=!0},p(t,s){const u={};s&4&&(u["aria-label"]=t[2]),e.$set(u)},i(t){l||(E(e.$$.fragment,t),l=!0)},o(t){N(e.$$.fragment,t),l=!1},d(t){te(e,t)}}}function Se(a){let e,l;return{c(){e=M("h3"),l=Q("Title")},l(t){e=I(t,"H3",{});var s=A(e);l=U(s,"Title"),s.forEach(g)},m(t,s){K(t,e,s),p(e,l)},p:be,d(t){t&&g(e)}}}function Ve(a){let e,l;return{c(){e=M("span"),l=Q("Content")},l(t){e=I(t,"SPAN",{});var s=A(e);l=U(s,"Content"),s.forEach(g)},m(t,s){K(t,e,s),p(e,l)},p:be,d(t){t&&g(e)}}}function ge(a){let e,l;return{c(){e=M("p"),l=Q(a[11]),this.h()},l(t){e=I(t,"P",{class:!0});var s=A(e);l=U(s,a[11]),s.forEach(g),this.h()},h(){D(e,"class","mt-2 text-sm font-normal text-danger")},m(t,s){K(t,e,s),p(e,l)},p(t,s){s&2048&&le(l,t[11])},d(t){t&&g(e)}}}function qe(a){let e;return{c(){e=Q(a[2])},l(l){e=U(l,a[2])},m(l,t){K(l,e,t)},p(l,t){t&4&&le(e,l[2])},d(l){l&&g(e)}}}function he(a){let e,l;return e=new ve({props:{thin:!0,variant:a[3],loading:a[6],disabled:a[4]||a[6],testId:"confirm-modal-button",type:"submit",$$slots:{default:[He]},$$scope:{ctx:a}}}),{c(){x(e.$$.fragment)},l(t){$(e.$$.fragment,t)},m(t,s){ee(e,t,s),l=!0},p(t,s){const u={};s&8&&(u.variant=t[3]),s&64&&(u.loading=t[6]),s&80&&(u.disabled=t[4]||t[6]),s&8388610&&(u.$$scope={dirty:s,ctx:t}),e.$set(u)},i(t){l||(E(e.$$.fragment,t),l=!0)},o(t){N(e.$$.fragment,t),l=!1},d(t){te(e,t)}}}function He(a){let e;return{c(){e=Q(a[1])},l(l){e=U(l,a[1])},m(l,t){K(l,e,t)},p(l,t){t&2&&le(e,l[1])},d(l){l&&g(e)}}}function Le(a){let e,l,t,s,u,b,m,H,F,L,v,k,O,S,V,c,q,R,f=!a[6]&&_e(a);const P=a[21].title,_=ae(P,a,a[23],me),h=_||Se(),j=a[21].content,B=ae(j,a,a[23],ue),T=B||Ve();let d=a[11]&&ge(a);k=new ve({props:{thin:!0,variant:"secondary",disabled:a[6],$$slots:{default:[qe]},$$scope:{ctx:a}}}),k.$on("click",a[12]);let r=!a[0]&&he(a),n=[{id:a[8]},{class:S="body "+a[9]},{"aria-modal":"true"},{"aria-labelledby":"modal-title"},{"data-testid":V=a[16]["data-testid"]},a[17]],y={};for(let i=0;i<n.length;i+=1)y=w(y,n[i]);return{c(){e=M("dialog"),f&&f.c(),l=z(),t=M("div"),h&&h.c(),u=z(),b=M("form"),m=M("div"),T&&T.c(),H=z(),d&&d.c(),L=z(),v=M("div"),x(k.$$.fragment),O=z(),r&&r.c(),this.h()},l(i){e=I(i,"DIALOG",{id:!0,class:!0,"aria-modal":!0,"aria-labelledby":!0,"data-testid":!0});var o=A(e);f&&f.l(o),l=J(o),t=I(o,"DIV",{id:!0,class:!0});var C=A(t);h&&h.l(C),C.forEach(g),u=J(o),b=I(o,"FORM",{method:!0});var W=A(b);m=I(W,"DIV",{id:!0,class:!0});var X=A(m);T&&T.l(X),H=J(X),d&&d.l(X),X.forEach(g),L=J(W),v=I(W,"DIV",{class:!0});var Y=A(v);$(k.$$.fragment,Y),O=J(Y),r&&r.l(Y),Y.forEach(g),W.forEach(g),o.forEach(g),this.h()},h(){D(t,"id",s="modal-title-"+a[8]),D(t,"class","title svelte-9mpp9v"),D(m,"id",F="modal-content-"+a[8]),D(m,"class","content svelte-9mpp9v"),D(v,"class","flex items-center justify-end space-x-2 p-6"),D(b,"method","dialog"),ie(e,y),G(e,"large",a[5]),G(e,"hightlightNav",a[7]),G(e,"svelte-9mpp9v",!0)},m(i,o){K(i,e,o),f&&f.m(e,null),p(e,l),p(e,t),h&&h.m(t,null),p(e,u),p(e,b),p(b,m),T&&T.m(m,null),p(m,H),d&&d.m(m,null),p(b,L),p(b,v),ee(k,v,null),p(v,O),r&&r.m(v,null),a[22](e),c=!0,q||(R=[Z(window,"click",a[15]),Z(window,"keydown",Ee(a[14])),Z(b,"submit",Ce(a[13]))],q=!0)},p(i,[o]){i[6]?f&&(ne(),N(f,1,1,()=>{f=null}),se()):f?(f.p(i,o),o&64&&E(f,1)):(f=_e(i),f.c(),E(f,1),f.m(e,l)),_&&_.p&&(!c||o&8388608)&&oe(_,P,i,i[23],c?re(P,i[23],o,Fe):fe(i[23]),me),(!c||o&256&&s!==(s="modal-title-"+i[8]))&&D(t,"id",s),B&&B.p&&(!c||o&8388608)&&oe(B,j,i,i[23],c?re(j,i[23],o,Ae):fe(i[23]),ue),i[11]?d?d.p(i,o):(d=ge(i),d.c(),d.m(m,null)):d&&(d.d(1),d=null),(!c||o&256&&F!==(F="modal-content-"+i[8]))&&D(m,"id",F);const C={};o&64&&(C.disabled=i[6]),o&8388612&&(C.$$scope={dirty:o,ctx:i}),k.$set(C),i[0]?r&&(ne(),N(r,1,1,()=>{r=null}),se()):r?(r.p(i,o),o&1&&E(r,1)):(r=he(i),r.c(),E(r,1),r.m(v,null)),ie(e,y=De(n,[(!c||o&256)&&{id:i[8]},(!c||o&512&&S!==(S="body "+i[9]))&&{class:S},{"aria-modal":"true"},{"aria-labelledby":"modal-title"},(!c||o&65536&&V!==(V=i[16]["data-testid"]))&&{"data-testid":V},o&131072&&i[17]])),G(e,"large",i[5]),G(e,"hightlightNav",i[7]),G(e,"svelte-9mpp9v",!0)},i(i){c||(E(f),E(h,i),E(T,i),E(k.$$.fragment,i),E(r),c=!0)},o(i){N(f),N(h,i),N(T,i),N(k.$$.fragment,i),N(r),c=!1},d(i){i&&g(e),f&&f.d(),h&&h.d(i),T&&T.d(i),d&&d.d(),te(k),r&&r.d(),a[22](null),q=!1,Ne(R)}}}function Oe(a,e,l){const t=["hideConfirm","confirmText","cancelText","confirmType","confirmDisabled","large","loading","hightlightNav","id","open","close","setError","class"];let s=ce(e,t),{$$slots:u={},$$scope:b}=e,{hideConfirm:m=!1}=e,{confirmText:H="Confirm"}=e,{cancelText:F="Cancel"}=e,{confirmType:L="primary"}=e,{confirmDisabled:v=!1}=e,{large:k=!1}=e,{loading:O=!1}=e,{hightlightNav:S=!1}=e,{id:V}=e,c;const q=()=>_.showModal(),R=()=>{l(11,c=""),_.close()},f=n=>{l(11,c=n)};let{class:P=""}=e,_;const h=ye(),j=()=>{R(),h("cancelModal")},B=()=>{l(11,c=""),h("confirmModal")},T=n=>{if(!q)return;const y=Array.from(_.querySelectorAll('button, input, div[contenteditable="true"]')).filter(C=>C instanceof HTMLDivElement?C.isContentEditable:!C.disabled),i=y[0],o=y[y.length-1];n.key==="Tab"&&(n.shiftKey?document.activeElement===i&&(o.focus(),n.preventDefault()):document.activeElement===o&&(i.focus(),n.preventDefault()))},d=n=>{n.target===_&&j()};function r(n){Me[n?"unshift":"push"](()=>{_=n,l(10,_)})}return a.$$set=n=>{l(16,e=w(w({},e),de(n))),l(17,s=ce(e,t)),"hideConfirm"in n&&l(0,m=n.hideConfirm),"confirmText"in n&&l(1,H=n.confirmText),"cancelText"in n&&l(2,F=n.cancelText),"confirmType"in n&&l(3,L=n.confirmType),"confirmDisabled"in n&&l(4,v=n.confirmDisabled),"large"in n&&l(5,k=n.large),"loading"in n&&l(6,O=n.loading),"hightlightNav"in n&&l(7,S=n.hightlightNav),"id"in n&&l(8,V=n.id),"class"in n&&l(9,P=n.class),"$$scope"in n&&l(23,b=n.$$scope)},a.$$.update=()=>{a.$$.dirty&1024&&q&&_&&_.focus()},e=de(e),[m,H,F,L,v,k,O,S,V,P,_,c,j,B,T,d,e,s,q,R,f,u,r,b]}class Ge extends ke{constructor(e){super(),Te(this,e,Oe,Le,pe,{hideConfirm:0,confirmText:1,cancelText:2,confirmType:3,confirmDisabled:4,large:5,loading:6,hightlightNav:7,id:8,open:18,close:19,setError:20,class:9})}get open(){return this.$$.ctx[18]}get close(){return this.$$.ctx[19]}get setError(){return this.$$.ctx[20]}}export{Ge as M};