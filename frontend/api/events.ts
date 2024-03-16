import client from "./client";

export async function getAllEvents() {
  const req = {
    limit: 5,
    offset: 0
  };

  const { data, error } = await client.GET("/events", {
    params: { query: req }
  });

  if (error) {
    console.error(error);
    return [];
  }

  return data;
}

export async function getEventById(uuid: string) {
  const { data, error } = await client.GET("/events/{id}", {
    params: {
      path: {
        id: uuid
      }
    }
  });
  if (error) {
    console.error(error);
    return [];
  }

  return data;
}
