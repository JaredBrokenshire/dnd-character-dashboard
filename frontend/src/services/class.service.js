import api from "@/services/api";
import HelperService from "@/services/helper.service";

const BASE_URL = "/classes"

const ClassService =  {
  async list(filters) {
    const queryString = HelperService.buildSearchQueryString(filters)
    return api.get(`${BASE_URL}${queryString}`)
  }
}

export default ClassService