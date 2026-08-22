# Generates the unknownFieldsModel test constant, a model carrying fields x/coremlcompiler
# does not model: Model.isUpdatable, ModelDescription.predictedFeatureName and
# .trainingInput, FeatureDescription.shortDescription and
# ArrayFeatureType.shapeRange. Written by coremltools, not by our encoder.
#
# Usage: python3 gen_unknownfields.py   (prints the hex for unknownFieldsModel)
import coremltools.proto.Model_pb2 as M

m = M.Model()
m.specificationVersion = 7
m.isUpdatable = True

d = m.description
d.metadata.shortDescription = "fixture"

inp = d.input.add()
inp.name = "x"
inp.shortDescription = "an input"
arr = inp.type.multiArrayType
arr.shape.append(4)
arr.dataType = M.ArrayFeatureType.FLOAT32
r = arr.shapeRange.sizeRanges.add()
r.lowerBound = 1
r.upperBound = 16

out = d.output.add()
out.name = "y"
out.type.multiArrayType.shape.append(4)
out.type.multiArrayType.dataType = M.ArrayFeatureType.FLOAT32

d.predictedFeatureName = "y"
ti = d.trainingInput.add()
ti.name = "x"
ti.type.multiArrayType.shape.append(4)
ti.type.multiArrayType.dataType = M.ArrayFeatureType.FLOAT32

m.mlProgram.version = 1

b = m.SerializeToString().hex()
for i in range(0, len(b), 64):
    print('\t"%s" +' % b[i:i+64])
