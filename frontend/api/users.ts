import client from "./client";
import type { paths } from "./schema";

export async function getUsers(queryParams?: paths["/users"]["get"]["parameters"]["query"]) {
  const { data, error } = await client.GET("/users", {
    params: { query: queryParams }
  });
  if (error) {
    return [];
  }
  return data;
}

export async function createUser(
  body: paths["/users"]["post"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.POST("/users", {
    body
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function deleteUser(pathParams: paths["/users/{id}"]["delete"]["parameters"]["path"]) {
  const { data, error } = await client.DELETE("/users/{id}", {
    params: { path: pathParams }
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function getUserById(pathParams: paths["/users/{id}"]["get"]["parameters"]["path"]) {
  const { data, error } = await client.GET("/users/{id}", {
    params: {
      path: pathParams
    }
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function updateUser(
  pathParams: paths["/users/{id}"]["patch"]["parameters"]["path"],
  body: paths["/users/{id}"]["patch"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.PATCH("/users/{id}", {
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

export async function saveUser(
  pathParams: paths["/users/{id}"]["patch"]["parameters"]["path"],
  body: paths["/users/{id}"]["put"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.PUT("/users/{id}", {
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

export default async function getMatchesByUserId(uuid: string) {
  const { data, error } = await client.GET("/matches/{id}", {
    params: {
      path: {
        id: uuid
      }
    }
  });

  if (error) {
    throw new Error("Failed to get user matches by ID");
  }

  return data;
}
