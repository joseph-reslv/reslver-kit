root=$(pwd)
source="${root}/sources/"

reslver="reslver/"
reslver_tf_loader="reslver-tf-loader/"
reslver_graph_generator="reslver-static-graph-exporter/" 
reslver_graph_exporter="reslver-graph-exporter/"

reslver_path="${root}/reslver/sources/"
reslver_tf_loader_path="${root}/reslver-tf-loader/sources/"
reslver_graph_generator_path="${root}/reslver-static-graph-exporter/reslver-graph/" 
reslver_graph_exporter_path="${root}/reslver-graph-exporter/sources/"

echo "Moving: Relsver..."
cp -r "${reslver_path}" "${source}${reslver}"

echo "Moving: Relsver TF Loader..."
cp -r "${reslver_tf_loader_path}" "${source}${reslver_tf_loader}"

echo "Building: Reslver Static Graph Generator"
python3 ./install.py
cd ${reslver_graph_generator_path}
pyinstaller --clean --onefile "reslvergraph.py" --distpath ${source}${reslver_graph_generator}
cd ${root}

echo "Moving: Relsver Graph Exporter..."
cp -r "${reslver_graph_exporter_path}" "${source}${reslver_graph_exporter}"

# point back to current dir
cd ${root}