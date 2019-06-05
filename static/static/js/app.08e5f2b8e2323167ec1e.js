webpackJsonp([1],{"4Y1N":function(t,e){},BKSd:function(t,e){},HFuo:function(t,e){},NHnr:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var r=a("7+uW"),s={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{attrs:{id:"app"}},[e("router-view")],1)},staticRenderFns:[]};var n=a("VU/8")({name:"App"},s,!1,function(t){a("Shb0")},null,null).exports,i=a("/ocq"),o={render:function(){this.$createElement;this._self._c;return this._m(0)},staticRenderFns:[function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"w-full border-b border-solid border-0 border-gray-200",attrs:{id:"nav"}},[e("ul",{staticClass:"flex flex-1 list-none m-0 pt-4 pb-4"},[e("a",{attrs:{href:"/products"}},[e("li",[this._v("PRODUCTS")])]),this._v(" "),e("a",{attrs:{href:"/categories"}},[e("li",[this._v("CATEGORIES")])])])])}]};var c=a("VU/8")(null,o,!1,function(t){a("4Y1N")},null,null).exports,l={name:"ProductImage",props:["product","image_id"],data:function(){return{image_data:""}},mounted:function(){var t=this;axios.get("http://localhost/api/products/image?id="+this.image_id).then(function(e){t.image_data=e.data})},methods:{openProduct:function(){this.$router.push("/product/"+this.product.UUID)}}},u={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"flex flex-1 h-full"},[a("div",{staticClass:"w-1/2 h-full overflow-hidden rounded-tl rounded-bl cursor-pointer",on:{click:t.openProduct}},[a("div",{staticClass:"w-full h-full flex flex-1 justify-center"},[a("img",{staticClass:"h-full mx-auto flex-none",attrs:{src:this.image_data}})])]),t._v(" "),a("div",{staticClass:"p-2"},[a("span",{staticClass:"cursor-pointer",on:{click:t.openProduct}},[t._v(t._s(t.product.Name))]),t._v(" "),a("div",{staticClass:"text-gray-400"},[t._v(t._s(t.product.Description))])])])},staticRenderFns:[]},d=a("VU/8")(l,u,!1,null,null,null).exports,p={name:"Index",mounted:function(){var t=this;axios.get("http://localhost/api/products/all").then(function(e){t.products=e.data})},data:function(){return{products:[]}},methods:{},components:{navbar:c,"product-tile":d}},h={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"w-full h-full"},[e("navbar"),this._v(" "),e("div",{staticClass:"flex flex-1 flex-wrap"},this._l(this.products,function(t,a){return e("div",{key:a,staticClass:"product my-2 h-64 border rounded border-solid border-0 hover:border-gray-400 border-gray-200"},[e("product-tile",{attrs:{product:t,image_id:JSON.parse(t.Images)[0]}})],1)}),0)],1)},staticRenderFns:[]};var f=a("VU/8")(p,h,!1,function(t){a("RC/g")},null,null).exports,m={render:function(){var t=this.$createElement;return(this._self._c||t)("div",{staticClass:"rounded bg-black text-white cursor-pointer p-4",on:{click:this.openCategory}},[this._v("\n    "+this._s(this.category.Name)+"\n")])},staticRenderFns:[]},g={name:"Index",data:function(){return{categories:[]}},components:{navbar:c,category:a("VU/8")({name:"CategoryTile",props:["category"],methods:{openCategory:function(){this.$router.push("/category/"+this.category.Name)}}},m,!1,null,null,null).exports},mounted:function(){var t=this;axios.get("http://localhost/api/categories/all").then(function(e){t.categories=e.data})}},v={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"w-full h-full"},[e("navbar"),this._v(" "),this._l(this.categories,function(t,a){return e("div",{key:a,staticClass:"category mt-2"},[e("category",{attrs:{category:t}})],1)})],2)},staticRenderFns:[]};var _=a("VU/8")(g,v,!1,function(t){a("BKSd")},null,null).exports,x={name:"CategoryPage",mounted:function(){var t=this;axios.get("http://localhost/api/categories/specific?name="+this.$route.params.name).then(function(e){t.products=e.data.Products})},data:function(){return{products:[]}},methods:{},components:{navbar:c,"product-tile":d}},C={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"w-full h-full"},[e("navbar"),this._v(" "),e("div",{staticClass:"flex flex-1 flex-wrap"},this._l(this.products,function(t,a){return e("div",{key:a,staticClass:"product my-2 h-64 border rounded border-solid border-0 hover:border-gray-400 border-gray-200"},[e("product-tile",{attrs:{product:t,image_id:JSON.parse(t.Images)[0]}})],1)}),0)],1)},staticRenderFns:[]};var b=a("VU/8")(x,C,!1,function(t){a("HFuo")},null,null).exports,w={name:"ProductImage",props:["src","load"],data:function(){return{image_data:"",image_set:!1,new_image:this.src}},mounted:function(){this.LoadImage()},methods:{SetImage:function(t){this.image_set=!0,this.new_image=t,this.LoadImage()},LoadImage:function(){var t=this;if(this.load){var e="http://localhost/api/products/image?id=";this.image_set?e+=this.new_image:e+=this.src,axios.get(e).then(function(e){t.image_data=e.data})}else requestIdleCallback(this.LoadImage)}}},y={render:function(){var t=this.$createElement;return(this._self._c||t)("img",{attrs:{id:this.new_image,src:this.image_data}})},staticRenderFns:[]},I={name:"Index",data:function(){return{product:{Images:[]},loaded:!1}},components:{navbar:c,"product-image":a("VU/8")(w,y,!1,null,null,null).exports},beforeCreate:function(){var t=this;axios.get("http://localhost/api/products/specific?id="+this.$route.params.uuid).then(function(e){e.data.Images=JSON.parse(e.data.Images),t.product=e.data,t.loaded=!0})},methods:{setPreviewImage:function(t){this.$refs.previewImage.SetImage(t.srcElement.id)}}},$={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"w-full h-full"},[a("navbar"),t._v(" "),a("div",{staticClass:"flex flex-1 w-full h-full"},[a("div",{staticClass:"w-1/2 h-full"},[a("product-image",{ref:"previewImage",staticClass:"w-full m-auto",attrs:{load:t.loaded,src:this.product.Images[0]}}),t._v(" "),a("ul",{staticClass:"flex flex-1 w-full"},t._l(this.product.Images,function(e,r){return a("li",{key:r,staticClass:"w-20 h-20 overflow-hidden",on:{click:t.setPreviewImage}},[a("product-image",{staticClass:"w-full cursor-pointer",attrs:{load:t.loaded,src:e}})],1)}),0)],1),t._v(" "),a("div",{staticClass:"w-1/2 p-4"},[a("h2",[t._v(t._s(t.product.Name))]),t._v(" "),a("div",[t._v(t._s(t.product.Description))])])])],1)},staticRenderFns:[]},E=a("VU/8")(I,$,!1,null,null,null).exports;r.a.use(i.a);var P=new i.a({mode:"history",routes:[{path:"/",name:"Index",component:f},{path:"/products",name:"Products",component:f},{path:"/categories",name:"Categories",component:_},{path:"/category/:name",name:"Category",component:b},{path:"/product/:uuid",name:"Product",component:E}]}),R=a("mtWM"),U=a.n(R);window.axios=U.a,r.a.config.productionTip=!1,new r.a({el:"#app",router:P,components:{App:n},template:"<App/>"})},"RC/g":function(t,e){},Shb0:function(t,e){}},["NHnr"]);
//# sourceMappingURL=app.08e5f2b8e2323167ec1e.js.map