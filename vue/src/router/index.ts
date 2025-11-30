import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Login from '../components/Login.vue'
import Library from '../components/Library.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/library',
    name: 'Library',
    component: Library
  }
]

export const router = createRouter({
    history: createWebHistory(),
    routes
});
