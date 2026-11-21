import BackgroundTexture from "@/assets/images/pages/background-texture.png"

export default [
  {
    path: "/classes",
    name: "class-list",
    component: () => import("@/views/classes/ClassList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  },
  {
    path: "/classes/:id",
    name: "class-details",
    component: () => import("@/views/classes/ClassDetails.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
      backgroundImage: BackgroundTexture,
    }
  }
]