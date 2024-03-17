import client from "./client";

export async function getAllEvents() {
  const req = {
    limit: 20,
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

export async function eventSwipe(userId: string, eventId: string, liked: boolean) {
  const { data, error } = await client.POST("/events/swipes", {
    body: {
      userId,
      eventId,
      liked
    }
  });

  if (error) {
    console.error(error);
    return [];
  }

  return data;
}
