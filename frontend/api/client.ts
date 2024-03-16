import createClient from "openapi-fetch";
import type { paths } from "./schema.d";

// export default createClient<paths>({ baseUrl: `http://localhost:8080` });

// this is my phone IP when using expo?? only way i could get it to work but we need to not hardcode
//  TODO: MAKE SURE THIS IS THE RIGHT IP
// (with expo running, shake your phone and find the IP. leave 8080 alone instead of 8081)
export default createClient<paths>({ baseUrl: `http://10.110.30.75:8080` });
