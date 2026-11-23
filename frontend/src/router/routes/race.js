export default [
  {
    path: "/races",
    name: "races-list-page",
    component: () => import("@/views/races/RaceList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  },
]