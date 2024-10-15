// src/pages/MainPage.jsx
import React, { useState } from "react";
import FileUpload from "../component/FileUpload";
import Map from "../component/Map";

const MainPage = () => {
  const [geojsonData, setGeojsonData] = useState(null);

  // Function to handle file upload
  const handleFileUpload = (fileContent) => {
    setGeojsonData(fileContent); // Set the GeoJSON or KML file content
  };

  return (
    <div>
      {/* File Upload Component */}
      <FileUpload onFileUpload={handleFileUpload} />

      {/* Map Component */}
      <Map geojsonData={geojsonData} />
    </div>
  );
};

export default MainPage;
