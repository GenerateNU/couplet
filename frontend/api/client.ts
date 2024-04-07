import createClient from "openapi-fetch";
import type { paths } from "./schema.d";

export default createClient<paths>({ baseUrl: `https://2980-155-33-135-28.ngrok-free.app` }); // set to your IP when running locally
