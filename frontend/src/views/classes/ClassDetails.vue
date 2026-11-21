<template>
  <div>
    <div
      v-if="classType"
      class="relative flex flex-col lg:flex-row gap-4 mt-8"
    >

      <nav class="lg:top-8 lg:sticky w-full h-fit transition-all bg-white lg:max-w-[25%]">
        <div class="flex flex-col border border-collapse transition-all">

          <!-- Back to top -->
          <button
            :class="`text-left px-4 transition-all overflow-hidden ${showBackToTop ? 'h-full py-2' : 'h-0'}`"
            type="button"
            @click="scrollTo('page-top')"
          >
            Back to Top
          </button>

          <!-- Class heading -->
          <h4 class="w-full font-bold">
            <button
              class="nav-button"
              type="button"
              @click="scrollTo('class-heading')"
            >
              {{ classType.name }}
            </button>
          </h4>

          <!-- Core traits -->
          <button
            :aria-expanded="coreTraitsOpen"
            aria-controls="core-traits-navigation"
            class="nav-button flex-between"
            type="button"
          >
            <span
              class="flex-grow"
              @click="scrollTo('core-traits')"
            >
              Core Traits
            </span>

            <feather-icon
              :class="`w-5 h-5 transition-all ${coreTraitsOpen ? 'rotate-90' : ''}`"
              icon="ChevronRightIcon"
              @click="coreTraitsOpen = !coreTraitsOpen"
            />
          </button>
          
          <button
            v-if="coreTraitsOpen"
            class="nav-button indented"
            type="button"
            @click="scrollTo('becoming-class')"
          >
            Becoming a{{ ['a','e','i','o','u'].includes(classType.name[0].toLowerCase()) ? 'n' : '' }} {{ classType.name }}
          </button>

          <!-- Class features -->
          <button
            class="nav-button"
            type="button"
            @click="scrollTo('class-features')"
          >
            Class Features
          </button>

          <!-- Subclasses -->

          <button
            class="nav-button"
            type="button"
            @click="scrollTo('class-subclasses')"
          >
            Subclasses
          </button>

          <button
            v-for="subclass in classType.subclasses"
            :key="`nav-button-subclass-${subclass.name.toLowerCase().replace(' ','_')}`"
            class="nav-button indented"
            type="button"
            @click="scrollTo(`subclass-${subclass.name.toLowerCase().replace(' ','_')}`)"
          >
            {{ subclass.name }}
          </button>

        </div>
      </nav>

      <div class="relative flex-grow">
        <div class="hidden xs:block float-right w-fit max-w-[40%] my-auto pb-[15px]">
          <img
            :src="classImage"
            alt=""
            class="block w-full max-w-[330px] my-auto"
          >
        </div>

        <div class="whitespace-pre-line">
          <h1
            id="class-heading"
            class="text-ink mb-4 md:mb-8"
          >
            {{ classType.name }}
          </h1>

          <section
            id="core-traits"
            class="overflow-x-auto mb-4"
          >
            <table>
              <caption>Core {{ classType.name }} Traits</caption>
              <thead />
              <tbody>
                <tr>
                  <th>Primary Ability</th>
                  <td>{{ classType.primary_ability }}</td>
                </tr>
                <tr>
                  <th>Hit Point Die</th>
                  <td>D{{ classType.hit_point_die_value }} per {{ classType.name }} level</td>
                </tr>
                <tr>
                  <th>Saving Throw Proficiencies</th>
                  <td>{{ classType.save_proficiencies.map(saveProf => saveProf.save).join(" & ") }}</td>
                </tr>
                <tr>
                  <th>Skill Proficiencies</th>
                  <td>
                    {{ classType.skill_proficiencies.length === 18
                      ?
                        `Choose any ${classType.skill_proficiency_max_options} skills`
                      :
                        `Choose ${classType.skill_proficiency_max_options}: ${classType.skill_proficiencies.map(skillProf => skillProf.skill).join(", ")}`
                    }}
                  </td>
                </tr>
                <tr>
                  <th>Weapon Proficiencies</th>
                  <td>{{ classType.weapon_proficiencies.map(weaponProf => `${weaponProf.weapon_type} weapons ${weaponProf.condition && `that have the ${weaponProf.condition} property`}`).join(" & ") }}</td>
                </tr>
                <tr>
                  <th>Armour Proficiencies</th>
                  <td>{{ armourProficienciesText }}</td>
                </tr>
              </tbody>
            </table>
          </section>

          <pre class="section-mb"><span>{{ classType.description }}</span></pre>

          <section
            id="becoming-class"
            class="section-mb"
          >
            <h2>Becoming a{{ ['a','e','i','o','u'].includes(classType.name[0].toLowerCase()) ? 'n' : '' }} {{ classType.name }}...</h2>
            <!-- Level 1 features -->
          </section>

          <section
            id="class-features"
            class="section-mb"
          >
            <h2>{{ classType.name }} Features</h2>
            <p>As a {{ classType.name }}, you gain the following class features when your reach the specified {{ classType.name }} levels. These features are listed in the {{ classType.name }} features table.</p>
          </section>

          <section
            id="class-subclasses"
            class="section-mb"
          >
            <h2>{{ classType.name }} Subclasses</h2>
            <p>A {{ classType.name }} subclass is a specialisation that grants you features at certain {{ classType.name }} levels, as specified in the subclass.</p>
          </section>

          <section
            v-for="subclass in classType.subclasses"
            :id="`subclass-${subclass.name.toLowerCase().replace(' ','_')}`"
            :key="`subclass-${subclass.name.toLowerCase().replace(' ','_')}`"
            class="section-mb"
          >
            <h2>{{ subclass.name }}</h2>
            <p>{{ subclass.description }}</p>
          </section>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
  import ClassService from "@/services/class.service";
  import HelperService from "@/services/helper.service";
  import FileService from "@/services/file.service";

  export default {
    name: "ClassDetails",
    data() {
      return {
        loading: true,
        classType: null,
        classImage: "",

        showBackToTop: false,
        coreTraitsOpen: false,
        subclassesOpen: false,
      }
    },
    computed: {
      armourProficienciesText() {
        const armourProficiencies = this.classType.armour_proficiencies
        if (!armourProficiencies) return "None"

        let shieldsIndex = null
        for (let i = 0; i < armourProficiencies.length; i++) {
          if (armourProficiencies[i].armour_type === "Shields") {
            shieldsIndex = i
            break
          }
        }

        if (shieldsIndex !== null) {
          if (armourProficiencies.length === 1) return "Shields"
          armourProficiencies.pop(shieldsIndex)
        }
        
        let text = ""
        for (let i = 0; i < armourProficiencies.length - 1; i++) {
          text += armourProficiencies[i].armour_type + ", "
        }

        text = text.slice(0, text.length - 2)

        return `${text.length > 0 ? `${text} and `: ''}${armourProficiencies[armourProficiencies.length-1].armour_type} armour${shieldsIndex !== null ? " and Shields" : ""}`
      }
    },
    mounted() {
      this.handleScroll()
      window.addEventListener("scroll", this.handleScroll, { passive: true })

      this.getClass().then(() => this.getImage())
    },
    beforeDestroy() {
      window.removeEventListener("scroll", this.handleScroll)
    },
    methods: {
      async getClass() {
        this.loading = true;

        try {
          const res = await ClassService.get(this.$route.params.id);
          this.classType = res.data;
          console.log(this.classType)
        } catch (error) {
          let errorText = "Could not get class, please refresh and try again";

          if (res && res.data.error) {
            errorText = res.data.error;
          }

          HelperService.errorToast(this.$toast, error, errorText);
        } finally {
          this.loading = false;
        }
      },
      async getImage() {
        const location = this.classType.details_image?.location;
        if (!location) return

        const res = await FileService.get(location);
        this.classImage = URL.createObjectURL(res.data);
      },
      handleScroll() {
        this.showBackToTop = window.scrollY > 200
      },

      scrollTo(id) {
        const element = document.getElementById(id)

        if (!element) return

        element.scrollIntoView({
          behavior: "smooth",
          block: "start",
        })
      },
    }
  }
</script>

<style scoped>
.section-mb {
  @apply mb-1 xs:mb-2 md:mb-4 xl:mb-8
}

h2 {
  @apply w-full mb-4 pb-4 border-b-2 border-gold
}

#class-heading,
#core-traits,
#becoming-class,
#class-features,
[id^="subclass-"] {
  scroll-margin-top: 100px;
}

.nav-button {
  @apply w-full min-h-fit text-left transition-all border px-4 py-2 relative
}

.nav-button::after {
  @apply content-[''] absolute top-0 left-0 w-[4px] h-full bg-transparent transition-all
}

.nav-button:hover::after {
  @apply bg-gold
}

.nav-button.indented {
  @apply pl-8
}
</style>