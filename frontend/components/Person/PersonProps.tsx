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
