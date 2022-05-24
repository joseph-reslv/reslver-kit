root=$(pwd)
source="${root}/sources/"

reslver="reslver/"
reslver_tf_loader="reslver-tf-loader/"
reslver_graph_generator="reslver-static-graph-exporter/" 
reslver_graph_exporter="reslver-graph-exporter/"
reslver_configs="reslver-configs/"

reslver_path="${root}/reslver/sources/"
reslver_tf_loader_path="${root}/reslver-tf-loader/sources/"
reslver_graph_generator_path="${root}/reslver-static-graph-exporter/reslver-graph/" 
reslver_graph_exporter_path="${root}/reslver-graph-exporter/sources/"
reslver_configs_path="${root}/reslver-configs/"

reslver_graph_generator_file="reslvergraph"

echo "Cloning: submodules..."
git submodule init
git submodule update --recursive --remote

echo "Copying: Relsver..."
cp -r "${reslver_path}" "${source}${reslver}"

echo "Copying: Relsver TF Loader..."
cp -r "${reslver_tf_loader_path}" "${source}${reslver_tf_loader}"

echo "Building: Reslver Static Graph Generator"
python3 ./install.py
cd ${reslver_graph_generator_path}
pyinstaller --clean --onefile "reslvergraph.py" --distpath ${source}${reslver_graph_generator} --log-level WARN
cd ${source}${reslver_graph_generator}
zip "${reslver_graph_generator_file}.zip" ${reslver_graph_generator_file}
rm ${reslver_graph_generator_file}
cd ${root}

echo "Copying: Relsver Graph Exporter..."
cp -r "${reslver_graph_exporter_path}" "${source}${reslver_graph_exporter}"

echo "Copying: Reslver Configs..."
cp -r "${reslver_configs_path}" "${source}${reslver_configs}"

# point back to current dir
cd ${root}

echo "Installing: Reslver Kit"
go install