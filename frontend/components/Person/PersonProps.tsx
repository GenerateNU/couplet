import type { components } from "../../api/schema";

type Event = components["schemas"]["Event"];

export type ImageWithCaptionProps = {
  image: string;
  caption?: string;
};

export type PersonHeightProps = {
  feet: number;
  inches: number;
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

  bio: string;
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
  mutualEvents: Event[];
  images: ImageWithCaptionProps[];
  isMatched: boolean;
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
}