<template>
  <div class="w-full mx-auto flex p-4 md:p-8">
    <div class="max-w-[1340px] mx-auto relative flex flex-grow justify-center md:items-center w-full min-h-full">
      <div class="flex-grow w-full h-full">
        <div class="w-full h-full relative overflow-hidden rounded-xl py-2 px-4 md:py-4 md:px-8">
          <div class="z-20 relative">
            <div class="flex-between">
              <h1 class="text-gold">MY CAMPAIGNS</h1>

              <button
                class="mt-4 lg:mt-2 bg-ember rounded-lg text-sm md:text-lg px-4 py-1 sm:px-4 sm:py-4 transition-all shadow-lg shadow-transparent text-gold-light hover:bg-ember-light hover:shadow-ember"
              >
                Start A Campaign
              </button>
            </div>

            <div class="w-full md:max-w-[50%] xl:max-w-[30%]">
              <input-wrapper
                v-model="filters.search"
                label="Search"
                name="campaigns-search"
                placeholder="Search"
                @input="onSearchInput"
              />
            </div>

            <div
              v-if="totalCount > 0 && !loading"
              class="flex-grow flex justify-center my-12 md:items-center"
            >
              <div class="h-fit w-full grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-12">
                <campaign-card
                  v-for="campaign in campaigns"
                  :key="`campaign-card-${campaign.id}`"
                  :campaign="campaign"
                  class="min-w-full h-full flex flex-col"
                />
              </div>

              <!-- TODO: Pagination -->
            </div>

            <div
              v-if="!loading && totalCount === 0"
              class="relative flex-grow flex justify-center md:items-center my-4 md:my-8 w-full h-fit"
            >

              <card
                :animated="false"
                class="w-[90%] xs:w-[75%] sm:w-[60%] lg:w-[40%] h-[80%]"
              >
                <template #content>
                  <div class="min-w-full min-h-full relative flex flex-col-reverse justify-between">
                    <div
                      class="flex flex-col justify-between items-start lg:justify-end gap-4"
                    >
                      <div class="sm:mx-2 xl:mx-8">
                        <div class="mb-4">
                          <h2 class="hidden sm:block text-gold uppercase">No Campaigns</h2>
                          <h4 class="block sm:hidden text-gold uppercase">No Campaigns</h4>
                        </div>
                        <p class="text-gold-light">
                          It doesn't look like you're in any campaigns yet. Start one now!
                        </p>
                      </div>

                      <div class="w-full flex-center">
                        <button
                          class="mt-4 lg:mt-2 bg-ember rounded text-sm md:text-lg px-2 py-1 sm:px-4 sm:py-2 transition-all shadow-lg shadow-transparent text-gold-light hover:bg-ember-light hover:shadow-ink"
                        >
                          Start a Campaign
                        </button>
                      </div>
                    </div>

                    <div class="min-w-[125px] min-h-[125px] hidden 2xs:flex justify-center items-center relative mx-auto w-[50%] lg:w-full 2xs:mb-8 lg:mb-16">
                      <img
                        alt=""
                        class="min-w-[50%] min-h-[50%] xs:max-w-[150%] xs:max-h-[150%] lg:max-w-[200%] lg:max-h-[200%] aspect-square object-contain absolute -top-[2rem] xs:-top-[3rem] sm:-top-[4rem] md:-top-[5rem] transition-all group-hover:scale-110"
                        src="@/assets/images/pages/character.webp"
                      >
                    </div>
                  </div>
                </template>
              </card>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import debounce from "lodash/debounce"
  import Card from "@/components/Card.vue"
  import HelperService from "@/services/helper.service";
  import CampaignService from "@/services/campaign.service";
  import InputWrapper from "@/components/forms/InputWrapper.vue";
  import CampaignCard from "@/views/campaigns/sections/CampaignCard.vue";

  export default {
    name: "CampaignsList",
    components: {CampaignCard, InputWrapper, Card},
    data() {
      return {
        loading: true,
        campaigns: [],
        totalCount: 0,
        filters: {
          page: 0,
          page_size: 12,
          search: "",
        },
        debouncedSearch: null,
      }
    },
    mounted() {
      this.debouncedSearch = debounce(this.getMoreCampaigns, 500)
      this.getCampaigns()
    },
    methods: {
      getMoreCampaigns(page) {
        this.filters.page = page;
        this.getCampaigns()
      },
      async getCampaigns() {
        this.loading = true;

        try {
          const res = await CampaignService.list(this.filters);

          this.campaigns = res.data.data || [];
          this.filters.page = res.data.meta.page
          this.filters.page_size = res.data.meta.page_size
          this.totalCount = res.data.meta.total_count
        } catch (error) {
          let errorText = "Could not get campaigns, please refresh and try again";

          if (res && res.data.error) {
            errorText = res.data.error;
          }

          HelperService.errorToast(this.$toast, errorText);
        } finally {
          this.loading = false;
        }
      },
      onSearchInput() {
        this.filters.page = 0
        this.debouncedSearch(0)
      }
    }
  }
</script>

<style scoped>

</style>