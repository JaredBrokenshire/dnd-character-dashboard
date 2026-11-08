<template>
  <component :is="layout">
    <div class="relative min-h-screen max-w-screen flex flex-col">
      <div class="w-full flex-grow mx-auto flex justify-center md:items-center p-4 md:p-8">
        <div class="max-w-[1340px] mx-auto relative flex-grow">
          <div class="flex-grow w-full h-full">
            <router-view class="flex-grow" />
          </div>
        </div>
      </div>
    </div>
  </component>
</template>

<script>
  const LayoutImageBackground = () => import("@/layouts/LayoutImageBackground.vue")

  export default {
    name: 'App',
    data() {
      return {
        layout: LayoutImageBackground,
      }
    },
  }
</script>
