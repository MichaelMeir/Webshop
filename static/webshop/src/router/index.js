import Vue from 'vue'
import Router from 'vue-router'
import Products from '@/components/Products'
import Categories from '@/components/Categories'

import CategoryPage from '@/components/CategoryPage'
import ProductPage from '@/components/ProductPage'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Products
    },
    {
      path: '/products',
      name: 'Products',
      component: Products
    },
    {
      path: '/categories',
      name: 'Categories',
      component: Categories
    },
    {
      path: '/category/:name',
      name: 'Category',
      component: CategoryPage
    },
    {
      path: '/product/:uuid',
      name: 'Product',
      component: ProductPage
    }
  ]
})
