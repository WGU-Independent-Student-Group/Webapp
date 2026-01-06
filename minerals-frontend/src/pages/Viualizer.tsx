import * as d3 from "d3";
import geoJSON from "../geoJSON/mockGeoJSON.json";

// Proof of concept to show d3 functioning, translating geo coordinates to display coordinates.
// Currently generating SVG paths, will probably want to use a Canvas
// Need to investigate GeoPermissableObject type to root out errors.
// Feature type will need to be tweaked if we need multiple properties
type Feature = {
  type: string;
  geometry: {
    type: string;
    coordinates: number[][];
  };
  properties: Record<string, string>;
};

type GeoJsonData = {
  type: string;
  features: Feature[];
};

type MapProps = {
  width: number;
  height: number;
};

const projection = d3.geoEquirectangular();

const geoPathGenerator = d3.geoPath().projection(projection);

export const Visualizer = ({ width, height }: MapProps) => {
  const data: GeoJsonData = geoJSON;
  const allSvgPaths = data.features.map((shape) => {
    return (
      <path
        key={shape.geometry.coordinates[0][0]} // lazy key, not garuanteed unique
        d={geoPathGenerator(shape)} // this errors because the current mock json is in an invalid shape for d3-geo
        stroke="black"
        fill="#cb1dd1"
      />
    );
  });

  return (
    <div>
      <svg width={width} height={height}>
        {allSvgPaths}
      </svg>
    </div>
  );
};
