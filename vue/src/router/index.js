import { createRouter, createWebHistory } from "vue-router";
import Login from "../components/Login.vue";
import Profile from "../components/Profile.vue";
import Library from "../components/Library.vue";
import Detailed from "../components/Detailed.vue";

const routes = [
  { path: "/", name: "Login", component: Login },
  { path: "/profile", name: "Profile", component: Profile },
  { path: "/library", name: "Library", component: Library },
  {
    path: "/album/:id",
    name: "AlbumDetail",
    component: Detailed,
    props: true, // pass route params if you want
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
