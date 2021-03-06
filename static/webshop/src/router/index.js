import Vue from 'vue'
import Router from 'vue-router'
import Products from '@/components/Products'
import Categories from '@/components/Categories'

import CategoryPage from '@/components/CategoryPage'
import ProductPage from '@/components/ProductPage'

import Auth from '@/components/Auth'

import AccountPanel from '@/components/AccountPanel'
import AdminPanel from '@/components/AdminPanel'

import ProcessOrder from '@/components/ProcessOrder'
import CompleteOrder from '@/components/CompleteOrder'

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
    },
    {
      path: '/auth',
      name: 'Authentication',
      component: Auth
    },
    {
      path: '/account',
      name: 'AccountPanel',
      component: AccountPanel
    },
    {
      path: '/admin',
      name: 'AdminPanel',
      component: AdminPanel
    },
    {
      path: '/processOrder',
      name: 'ProcessOrder',
      component: ProcessOrder
    },
    {
      path: '/completeOrder',
      name: 'CompleteOrder',
      component: CompleteOrder
    }
  ]
})
