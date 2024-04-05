import createClient from "openapi-fetch";
import type { paths } from "./schema.d";

export default createClient<paths>({ baseUrl: `http://10.110.207.81:8080` }); // set to your IP when running locally
