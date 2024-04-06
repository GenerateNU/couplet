import client from "./client";
import type { paths } from "./schema";

export async function getOrgs(queryParams?: paths["/orgs"]["get"]["parameters"]["query"]) {
  const { data, error } = await client.GET("/orgs", {
    params: { query: queryParams }
  });
  if (error) {
    return [];
  }
  return data;
}

export async function createOrg(
  body: paths["/orgs"]["post"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.POST("/orgs", {
    body
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function deleteOrg(pathParams: paths["/orgs/{id}"]["delete"]["parameters"]["path"]) {
  const { data, error } = await client.DELETE("/orgs/{id}", {
    params: { path: pathParams }
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function getOrgById(pathParams: paths["/orgs/{id}"]["get"]["parameters"]["path"]) {
  const { data, error } = await client.GET("/orgs/{id}", {
    params: {
      path: pathParams
    }
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function updateOrg(
  pathParams: paths["/orgs/{id}"]["patch"]["parameters"]["path"],
  body: paths["/orgs/{id}"]["patch"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.PATCH("/orgs/{id}", {
    params: {
      path: pathParams
    },
    body
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function saveOrg(
  pathParams: paths["/orgs/{id}"]["patch"]["parameters"]["path"],
  body: paths["/orgs/{id}"]["put"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.PUT("/orgs/{id}", {
    params: {
      path: pathParams
    },
    body
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}
