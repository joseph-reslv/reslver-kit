root=$(pwd)
source="${root}/sources/"

reslver="reslver/"
reslver_tf_loader="reslver-tf-loader/"
# reslver_excel_exporter="reslver-excel-exporter/" # NO EXCEL currently
reslver_graph_exporter="reslver-graph-exporter/"

reslver_path="${root}/reslver/sources/"
reslver_tf_loader_path="${root}/reslver-tf-loader/sources/"
# reslver_excel_exporter_path="${root}/reslver-excel-exporter/sources/" # NO EXCEL currently
reslver_graph_exporter_path="${root}/reslver-graph-exporter/sources/"

# build reslver
echo "Moving: Relsver..."
cp -r "${reslver_path}" "${source}${reslver}"

# build reslver
echo "Moving: Relsver TF Loader..."
cp -r "${reslver_tf_loader_path}" "${source}${reslver_tf_loader}"

# # build reslver # NO EXCEL currently
# echo "Moving: Relsver Excel Exporter..."
# cp -r "${reslver_excel_exporter_path}" "${source}${reslver_excel_exporter}"

# build reslver
echo "Moving: Relsver Graph Exporter..."
cp -r "${reslver_graph_exporter_path}" "${source}${reslver_graph_exporter}"

# point back to current dir
cd ${root}