export default [
  {
    path: "/classes",
    name: "classes-list",
    component: () => import("@/views/classes/ClassList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  }
]