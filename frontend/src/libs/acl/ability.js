import store from "@/store"
import {Ability} from "@casl/ability"
import {initialAbility} from "./config"

let userData = null
if (store.state.user.access_token) {
  userData = store.state.user.access_token
}

const existingAbility = userData ? userData.permissions : null

export default new Ability(existingAbility || initialAbility)
