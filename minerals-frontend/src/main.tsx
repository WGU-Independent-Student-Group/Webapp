import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { BrowserRouter, Routes, Route } from "react-router";
import { Home } from "./pages/Home.tsx";
import { Visualizer } from "./pages/Viualizer.tsx";
import { Commodity } from "./pages/Commodity.tsx";
import { About } from "./pages/About.tsx";
import { CommodityList } from "./pages/CommodityList.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="about" element={<About />} />
        <Route path="map" element={<Visualizer width={800} height={400} />} />
        <Route path="commodity">
          <Route index element={<CommodityList />} />
          <Route path=":commodity" element={<Commodity />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
