package main

import "testing"

func BenchmarkSerializeToJSON(b *testing.B){
    for i := 0; i < b.N; i++{
        serializeToJson(metadata)
    }
}

func BenchmarkSerializeToXML(b *testing.B){
    for i:=0; i<b.N; i++ {
        serializeToXML(metadata)
    }
}

func BenchmarkSerializeToProto(b *testing.B){
    for i:=0; i<b.N; i++{
        serializeToProto(genMetadata)
    }
}
