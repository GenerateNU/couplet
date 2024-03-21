import createClient from "openapi-fetch";
import type { paths } from "./schema.d";

// CHANGE TO URL BASE URL
export default createClient<paths>({ baseUrl: "http://10.110.63.100:8080" });
