<template>
  <router-link
    :to="`/campaigns/${$props.campaign.id}`"
    class=""
  >
    <card>
      <template #content>
        <div class="absolute -top-[1.5rem] left-0 right-0 w-full h-fit flex-center gap-1 z-20 transition-all group-hover:scale-110">
          <div
            v-for="(_, index) in $props.campaign.characters?.slice(0,5)"
            :key="`campaign-card-${$props.campaign.id}-character-profile-${index + 1}`"
            class="w-12 h-12 aspect-square rounded relative"
          >
            <img
              v-if="profilePictures[index]"
              :src="profilePictures[index]"
              alt="Profile Picture"
              class="absolute top-0 left-0 w-full h-full object-cover rounded"
            >
            <img
              v-else
              alt="Default Profile Picture"
              src="@/assets/images/pages/characters/default-character-profile.png"
            >
          </div>

          <div
            v-if="$props.campaign.characters?.length > 6"
            class="w-12 h-12 flex-center aspect-square rounded bg-ink border-2 border-gold"
          >
            <feather-icon
              class="w-6 h-6 text-gold"
              icon="PlusIcon"
            />
          </div>

        </div>


        <div class="w-full relative z-20 flex flex-grow gap-4 items-center">
          <div class="flex-center flex-col w-full gap-2 lg:gap-4 text-center">
            <h3 class="text-gold">{{ $props.campaign.name }}</h3>
            <h6 class="text-gold-light">Started {{ new Date($props.campaign.created_at).toLocaleDateString() }}</h6>
            <div class="flex-center flex-col">
              <h3 class="text-white">{{ $props.campaign.characters?.length || 0 }}</h3>
              <p class="text-sm text-white">Player{{ $props.campaign.characters?.length !== 1 ? 's' : '' }}</p>
            </div>
            <div class="flex-center flex-col">
              <p class="text-white">
                Role:
                {{ $props.campaign.dm.id === $store.getters["auth/currentUser"].id ? 'Dungeon Master' : 'Player' }}
              </p>
              <p
                v-if="$props.campaign.dm.id !== $store.getters['auth/currentUser'].id"
                class="text-white"
              >
                DM: {{ $props.campaign.dm.username }}</p>
            </div>
          </div>

        </div>
      </template>
    </card>
  </router-link>
</template>

<script>
  import Card from "@/components/Card.vue";
  import FileService from "@/services/file.service";

  export default {
    name: "CampaignCard",
    components: {Card},
    props: {
      campaign: {
        type: Object,
        required: true,
      },
    },
    data() {
      return {
        profilePictures: [],
      }
    },
    beforeUnmount() {
      for (let i = 0; i < this.profilePictures.length; i++) {
        URL.revokeObjectURL(this.profilePictures[i]);
      }
    },
    mounted() {
      this.getProfilePictures()
    },
    methods: {
      async getProfilePictures() {
        if (this.$props.campaign.characters == null) return
        
        for (let i=0; i < Math.min(this.$props.campaign.characters.length, 7); i++) {
          const character = this.$props.campaign.characters[i]
          try {
            const res = await FileService.get(character.profile_picture.location)
            this.profilePictures.push(URL.createObjectURL(res.data))
          } catch (e) {
            // No profile picture for character
            this.profilePictures.push("@/assets/images/pages/characters/default-character-profile.png")
          }

        }
      }
    }
  }
</script>

<style scoped>

</style>