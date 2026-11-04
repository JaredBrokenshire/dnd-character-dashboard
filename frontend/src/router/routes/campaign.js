export default [
  {
    path: "/campaigns",
    name: "Campaigns List",
    component: () => import("@/views/campaigns/CampaignsList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  }
]