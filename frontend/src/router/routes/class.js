export default [
  {
    path: "/classes",
    name: "Classes List",
    component: () => import("@/views/classes/ClassList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  }
]