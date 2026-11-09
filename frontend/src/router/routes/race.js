export default [
  {
    path: "/races",
    name: "Races List",
    component: () => import("@/views/races/RaceList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  },
]