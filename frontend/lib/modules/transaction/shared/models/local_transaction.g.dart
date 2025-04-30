// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'local_transaction.dart';

// **************************************************************************
// TypeAdapterGenerator
// **************************************************************************

class LocalTransactionAdapter extends TypeAdapter<LocalTransaction> {
  @override
  final int typeId = 0;

  @override
  LocalTransaction read(BinaryReader reader) {
    final numOfFields = reader.readByte();
    final fields = <int, dynamic>{
      for (int i = 0; i < numOfFields; i++) reader.readByte(): reader.read(),
    };
    return LocalTransaction(
      id: fields[0] as String,
      description: fields[1] as String,
      date: fields[2] as DateTime,
      amountUsd: fields[3] as double,
    );
  }

  @override
  void write(BinaryWriter writer, LocalTransaction obj) {
    writer
      ..writeByte(4)
      ..writeByte(0)
      ..write(obj.id)
      ..writeByte(1)
      ..write(obj.description)
      ..writeByte(2)
      ..write(obj.date)
      ..writeByte(3)
      ..write(obj.amountUsd);
  }

  @override
  int get hashCode => typeId.hashCode;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is LocalTransactionAdapter &&
          runtimeType == other.runtimeType &&
          typeId == other.typeId;
}
