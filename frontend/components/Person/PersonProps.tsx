export type ImageWithCaptionProps = {
  image: string;
  caption?: string;
};

export type PersonHeightProps = {
  feet: number;
  inches: number;
};

export type EventCardItemProps = {
  title: string;
  description: string;
  imageUrl: string;
};

export type PersonProps = {
  // Person info
  id: number;
  firstName: string;
  lastName: string;
  age: number;
  pronouns?: string;

  // Basic info
  location?: string;
  school?: string;
  work?: string;
  height?: PersonHeightProps;

  promptQuestion: string;
  promptResponse: string;
  interests: string[];

  // Lifestyle info
  relationshipType?: string;
  religion?: string;
  politicalAffiliation?: string;
  alchoholFrequency?: string;
  smokingFrequency?: string;
  drugFrequency?: string;
  cannabisFrequency?: string;

  instagramUsername: string;
  mutualEvents: EventCardItemProps[];
  images: ImageWithCaptionProps[];
  isMatched: boolean;
  likesYou: boolean;
  handleReact: (like: boolean) => void;
};

export type PillProps = {
  textColor: string;
  backgroundColor: string;
  items: string[];
};

export type LifestyleProps = {
  relationshipType?: string;
  religion?: string;
  politicalAffiliation?: string;
  alchoholFrequency?: string;
  smokingFrequency?: string;
  drugFrequency?: string;
  cannabisFrequency?: string;
};
