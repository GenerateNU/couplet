import client from "./client";
import type { paths } from "./schema";

export async function getEvents(queryParams?: paths["/events"]["get"]["parameters"]["query"]) {
  const { data, error } = await client.GET("/events", {
    params: { query: queryParams }
  });
  if (error) {
    return [];
  }
  return data;
}

export async function createEvent(
  body: paths["/events"]["post"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.POST("/events", {
    body
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function deleteEvent(
  pathParams: paths["/events/{id}"]["delete"]["parameters"]["path"]
) {
  const { data, error } = await client.DELETE("/events/{id}", {
    params: { path: pathParams }
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function getEventById(pathParams: paths["/events/{id}"]["get"]["parameters"]["path"]) {
  const { data, error } = await client.GET("/events/{id}", {
    params: {
      path: pathParams
    }
  });
  if (error) {
    throw new Error(`${error.code}: "${error.message}`);
  }
  return data;
}

export async function updateEvent(
  pathParams: paths["/events/{id}"]["patch"]["parameters"]["path"],
  body: paths["/events/{id}"]["patch"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.PATCH("/events/{id}", {
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

export async function saveEvent(
  pathParams: paths["/events/{id}"]["patch"]["parameters"]["path"],
  body: paths["/events/{id}"]["put"]["requestBody"]["content"]["application/json"]
) {
  const { data, error } = await client.PUT("/events/{id}", {
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

export async function eventSwipe(userId: string, eventId: string, liked: boolean) {
  const { data, error } = await client.POST("/events/swipes", {
    body: {
      userId,
      eventId,
      liked
    }
  });
  if (error) {
    throw new Error("Failed to swipe event");
  }
  return data;
}
