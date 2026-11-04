<template>
  <component :is="layout">
    <div
      v-if="showHeader"
      class="relative min-h-screen max-w-screen flex flex-col"
    >
      <page-header
        class="w-screen shadow-lg shadow-ink"
      />

      <router-view class="flex-grow" />
    </div>

    <router-view
      v-else
      class="relative min-h-screen"
    />
  </component>
</template>

<script>
  import PageHeader from "@/components/PageHeader.vue"

  const LayoutImageBackground = () => import("@/layouts/LayoutImageBackground.vue")

  export default {
    name: 'App',
    components: {
      PageHeader
    },
    data() {
      return {
        layout: LayoutImageBackground,
        showHeader: !!this.$route.meta.showHeader
      }
    },
    watch: {
      $route(to, _) {
        this.showHeader = !!to.meta.showHeader
      }
    }
  }
</script>
