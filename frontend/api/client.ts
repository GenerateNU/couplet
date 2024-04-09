import createClient from "openapi-fetch";
import type { paths } from "./schema.d";

export default createClient<paths>({ baseUrl: `${process.env.REACT_APP_API}` });