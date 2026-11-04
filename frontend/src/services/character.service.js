import api from "@/services/api";
import HelperService from "@/services/helper.service";

const BASE_URL = "/characters"

const CharacterService = {
  async list(filters) {
    const queryString = HelperService.buildSearchQueryString(filters)
    return api.get(`${BASE_URL}${queryString}`)
  },
  async get(id) {
    return api.get(`${BASE_URL}/${id}`)
  }
}

export default CharacterService