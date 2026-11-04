export default [
  {
    path: "/characters",
    name: "Characters List",
    component: () => import("@/views/characters/CharacterList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  }
]