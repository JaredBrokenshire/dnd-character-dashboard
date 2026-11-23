import api from "@/services/api";
import HelperService from "@/services/helper.service";

const BASE_URL = "/races"

const RaceService = {
  async list(filters) {
    const queryString = HelperService.buildSearchQueryString(filters)
    return api.get(`${BASE_URL}${queryString}`)
  },
}

export default RaceService;