import{S as m,b as p,a as u,C as c,D as f,E as _,j as g,g as l,F as $,K as E}from"../chunks/index.c37b9eda.js";/* empty css                    */import{p as d}from"../chunks/stores.5c33997d.js";import{E as h}from"../chunks/error.f78f02d8.js";import{i as b}from"../chunks/is-network-error.3dd5226d.js";import{p as y}from"../chunks/parse-with-big-int.02af2f95.js";function C(e){let t,o;return t=new h({props:{error:e[1],status:e[0]}}),{c(){c(t.$$.fragment)},l(r){f(t.$$.fragment,r)},m(r,s){_(t,r,s),o=!0},p(r,[s]){const a={};s&2&&(a.error=r[1]),s&1&&(a.status=r[0]),t.$set(a)},i(r){o||(g(t.$$.fragment,r),o=!0)},o(r){l(t.$$.fragment,r),o=!1},d(r){$(t,r)}}}function q(e,t,o){let r,s,a;E(e,d,i=>o(2,a=i));let n;try{n=y(r.message),b(n)&&(s=n.statusCode)}catch{}return e.$$.update=()=>{e.$$.dirty&4&&o(1,r=a.error),e.$$.dirty&4&&o(0,s=a.status)},[s,r,a]}class A extends m{constructor(t){super(),p(this,t,q,C,u,{})}}export{A as default};
