import esbuild from "esbuild";

esbuild
    .build({
        entryPoints: [
          "frontend/Application.jsx", 
          "frontend/style.css"
        ],
        outdir: "public/assets",
        bundle: true,
        plugins: [],
    })
    .then(() => console.log("⚡ Build complete! ⚡"))
    .catch(() => process.exit(1));
