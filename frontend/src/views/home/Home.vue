<template>
  <div class="w-full max-w-[1340px] mx-auto flex px-4 py-8 md:px-8 md:py-16">
    <div class="h-fit w-full grid grid-cols-1 lg:grid-cols-3 gap-8 md:gap-16">
      <home-card
        v-for="card in cards"
        :key="`home-card-${card.title.toLowerCase()}`"
        :body-text="card.bodyText"
        :image-path="card.imagePath"
        :route="card.route"
        :title="card.title"
      />
    </div>
  </div>
</template>

<script>
  import HomeCard from "@/views/home/sections/HomeCard.vue";
  import RacesImage from "@/assets/images/pages/home/races.png"
  import ClassesImage from "@/assets/images/pages/home/classes.png"
  import CampaignsImage from "@/assets/images/pages/home/campaigns.png"
  import CharactersImage from "@/assets/images/pages/home/characters.png"

  export default {
    name: "Home",
    components: {HomeCard},
    data() {
      return {
        cards: [
          {
            title: "Characters",
            bodyText: "Create your character in seconds with our digital builder",
            route: "/characters",
            imagePath: CharactersImage,
          },
          {
            title: "Campaigns",
            bodyText: "Collect all of your players in one place, keep track of rolls, and adventure together",
            route: "/campaigns",
            imagePath: CampaignsImage,
          },
          {
            title: "Classes",
            bodyText: "Choose from a wide range of classes, each provides a character’s most exciting capabilities",
            route: "/classes",
            imagePath: ClassesImage,
          },
          {
            title: "Races",
            bodyText: "Become a member of a fantastical species, and use their unique abilities",
            route: "/races",
            imagePath: RacesImage,
          }
        ]
      }
    },
  }
</script>

<style scoped />