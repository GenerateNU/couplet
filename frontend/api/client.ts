import createClient from "openapi-fetch";
import type { paths } from "./schema.d";

export default createClient<paths>({ baseUrl: `http://10.110.252.207:8080` });
