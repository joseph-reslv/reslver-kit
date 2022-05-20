root=$(pwd)
source="${root}/source/"

reslver="reslver"
reslver_tf_loader="reslver-tf-loader"
reslver_excel_exporter="reslver-excel-exporter"
reslver_graph_exporter="reslver-graph-exporter"

reslver_path="${root}/reslver"
reslver_tf_loader_path="${root}/reslver-tf-loader"
reslver_excel_exporter_path="${root}/reslver-excel-exporter"
reslver_graph_exporter_path="${root}/reslver-graph-exporter"

# build reslver
echo "Building: Relsver..."
cd ${reslver_path} && go build 
mv "${reslver_path}/${reslver}" "${source}/${reslver}"
echo "Built: Reslver"

# build reslver
echo "Building: Relsver TF Loader..."
cd ${reslver_tf_loader_path} && go build 
mv "${reslver_tf_loader_path}/${reslver_tf_loader}" "${source}/${reslver_tf_loader}"
echo "Built: Reslver TF Loader"

# build reslver
echo "Building: Relsver Excel Exporter..."
cd ${reslver_excel_exporter_path} && go build 
mv "${reslver_excel_exporter_path}/${reslver_excel_exporter}" "${source}/${reslver_excel_exporter}"
echo "Built: Reslver Excel Exporter"

# build reslver
echo "Building: Relsver Graph Exporter..."
cd ${reslver_graph_exporter_path} && go build 
mv "${reslver_graph_exporter_path}/${reslver_graph_exporter}" "${source}/${reslver_graph_exporter}"
echo "Built: Reslver Graph Exporter"

# point back to current dir
cd ${root}