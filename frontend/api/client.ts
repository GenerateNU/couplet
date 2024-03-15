import createClient from "openapi-fetch";
import type { paths } from "./schema.d";

// export default createClient<paths>({ baseUrl: `http://localhost:8080` });

// this is my phone IP when using expo?? only way i could get it to work but we need to not hardcode
//  TODO: fix this
export default createClient<paths>({ baseUrl: `http://10.110.150.59:8080` });
// export default createClient<paths>({ baseUrl: "http://192.68.1.12:8080" });
