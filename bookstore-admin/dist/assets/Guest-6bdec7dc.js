import{s as P,d as w,y as E,i,o as k,e as x,f as l,w as o,V as z,L as D,P as L,Q as N,r as T,c as B,E as F,b as _,h as p,M as d,C as b,W as v,D as s,m as I,O as V,R as C,X as M,S as R}from"./index-9e848275.js";/* empty css             *//* empty css                     *//* empty css                 *//* empty css              *//* empty css               *//* empty css                  */const $={__name:"Table",props:{action:String},setup(f,{expose:a}){const g=f,m=P(),{goTop:u}=m.appContext.config.globalProperties,e=w({loading:!1,tableData:[],total:0,currentPage:1,pageSize:10,multipleSelection:[]});E(()=>{n()});const n=()=>{e.loading=!0,i.get(g.action,{params:{pageNumber:e.currentPage,pageSize:e.pageSize}}).then(t=>{e.tableData=t.list,e.total=t.totalCount,e.currentPage=t.currPage,e.loading=!1,u&&u()})},h=t=>{e.multipleSelection=t},c=t=>{e.currentPage=t,n()};return a({state:e,getList:n}),(t,S)=>{const r=L,y=N;return k(),x(D,null,[l(r,{load:e.loading,data:e.tableData,"tooltip-effect":"dark",style:{width:"100%"},onSelectionChange:h},{default:o(()=>[z(t.$slots,"column")]),_:3},8,["load","data"]),l(y,{background:"",layout:"prev, pager, next",total:e.total,"page-size":e.pageSize,"current-page":e.currentPage,onCurrentChange:c},null,8,["total","page-size","current-page"])],64)}}},A={class:"header"},H={__name:"Guest",setup(f){let a=T(null);const g=()=>{if(!a.value.state.multipleSelection.length){s.error("请选择项");return}i.put("/users/0",{ids:a.value.state.multipleSelection.map(n=>n.userId)}).then(()=>{s.success("解除成功"),a.value.getList()})},m=()=>{if(!a.value.state.multipleSelection.length){s.error("请选择项");return}i.put("/users/1",{ids:a.value.state.multipleSelection.map(n=>n.userId)}).then(()=>{s.success("禁用成功"),a.value.getList()})},u=()=>{if(!a.value.state.multipleSelection.length){s.error("请选择项");return}i.put("/users/del/1",{ids:a.value.state.multipleSelection.map(n=>n.userId)}).then(()=>{s.success("删除成功"),a.value.getList()})},e=()=>{if(!a.value.state.multipleSelection.length){s.error("请选择项");return}i.put("/users/del/0",{ids:a.value.state.multipleSelection.map(n=>n.userId)}).then(()=>{s.success("恢复成功"),a.value.getList()})};return(n,h)=>{const c=I,t=V,S=F;return k(),B(S,{class:"guest-container"},{header:o(()=>[_("div",A,[l(c,{type:"primary",icon:d(C),onClick:g},{default:o(()=>[p("解除禁用")]),_:1},8,["icon"]),l(c,{type:"info",icon:d(M),onClick:m},{default:o(()=>[p("禁用账户")]),_:1},8,["icon"]),l(c,{type:"danger",icon:d(R),onClick:u},{default:o(()=>[p("删除账户")]),_:1},8,["icon"]),l(c,{type:"success",icon:d(C),onClick:e},{default:o(()=>[p("恢复账户")]),_:1},8,["icon"])])]),default:o(()=>[l($,{action:"/users",ref_key:"table",ref:a},{column:o(()=>[l(t,{type:"selection",width:"55"}),l(t,{prop:"nickName",label:"昵称"}),l(t,{prop:"loginName",label:"登录名"}),l(t,{label:"身份状态"},{default:o(r=>[_("span",{style:v(r.row.lockedFlag===0?"color: green;":"color: red;")},b(r.row.lockedFlag===0?"正常":"禁用"),5)]),_:1}),l(t,{label:"是否注销"},{default:o(r=>[_("span",{style:v(r.row.lockedFlag===0?"color: green;":"color: red;")},b(r.row.isDeleted===0?"正常":"注销"),5)]),_:1}),l(t,{prop:"createTime",label:"注册时间"})]),_:1},512)]),_:1})}}};export{H as default};