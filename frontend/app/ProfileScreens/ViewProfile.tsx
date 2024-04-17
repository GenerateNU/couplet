import { getEvents } from "../../api/events";
import { getUsers } from "../../api/users";
import type { EventCardItemProps, PersonProps } from "../../components/Person/PersonProps";
import { useAppSelector } from "../../state/hooks";
import calculateAge from "../../utils/calculateAge";

type User = Awaited<ReturnType<typeof getUsers>>[number];
type Event = Awaited<ReturnType<typeof getEvents>>[number];

export default function ViewProfile() {
  const userState = useAppSelector((state) => {
    return state.form;
  });
  const events: EventCardItemProps[] = [];
  getEvents({ limit: 4, offset: 0 }).then((fetchedEvents: Event[]) => {
    fetchedEvents.forEach((fetchedEvent: Event) => {
      events.push({
        title: fetchedEvent.name,
        description: fetchedEvent.bio,
        imageUrl: fetchedEvent.images[0]
      });
    });
  });
  const user: PersonProps = {
    id: userState.id,
    firstName: userState.fullName,
    lastName: "No Last Name",
    age: calculateAge(new Date(userState.birthday)),
    promptQuestion: userState.promptBio,
    promptResponse: userState.responseBio,
    interests: userState.passion,
    instagramUsername: userState.instagram,
    mutualEvents: events,
    images: userState.photos.map((photo) => {
      return { image: photo.filePath, caption: photo.caption };
    }),
    isMatched: true,
    likesYou: false,
    handleReact: () => {}
  };
}
