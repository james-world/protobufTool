using ProtoBuf;

var message = new ScheduleInstructionStep3V2
{
    CorrelationId = "123456",
    RequestId = Guid.Parse("30f6432c-ae25-4fbf-bd80-3b537e8ef0f2"),
    TenantId = Guid.Parse("a70d4280-7117-4fcc-a7a1-a185e465fddc"),
    CustomerRef = "MyCustomerRef",
    FinancialInstrumentId = Guid.Parse("a8f9f2f8-36f3-40b1-871a-70f92feb4c44"),
    WorkflowId = "MyWorkflowId",
    OrderRef = "MyOrderRef",
    InstructionRef = "MyInstructionRef",
    Country = "UK",
    Currency = "GBP",
    Amount = 10m,
    RequestedSettlementDate = DateTime.Parse("2023-08-01"),
    ForecastedSettlementDate = DateTime.Parse("2023-08-03"),
    ScheduledExecutionTimestamp = DateTime.Parse("2023-07-28"),
    AppId = "MyAppId",
    AppQueue = "Braintree",
    PaymentMethod = "CreditCard",
};

using (var file = File.Create("message.bin"))
{
	Serializer.Serialize(file, message);
};

using (var file = File.Create("message.txt"))
{
    var ms = new MemoryStream();
    Serializer.Serialize(ms, message);
    var b64 = Convert.ToBase64String(ms.ToArray());
    using var sw = new StreamWriter(file);
    sw.Write(b64);
}


[ProtoContract]
public class ScheduleInstructionStep3V2
{
    [ProtoMember(1)]
    public string CorrelationId { get; set; }

    [ProtoMember(2)]
    public Guid RequestId { get; set; }

    [ProtoMember(3)]
    public Guid TenantId { get; set; }

    [ProtoMember(4)]
    public string CustomerRef { get; set; }

    [ProtoMember(5)]
    public Guid FinancialInstrumentId { get; set; }

    [ProtoMember(6)]
    public string WorkflowId { get; set; }

    [ProtoMember(7)]
    public string OrderRef { get; set; }

    [ProtoMember(8)]
    public string InstructionRef { get; set; }

    [ProtoMember(9)]
    public string Country { get; set; }

    [ProtoMember(10)]
    public string Currency { get; set; }

    [ProtoMember(11)]
    public decimal Amount { get; set; }

    [ProtoMember(12)]
    public DateTime RequestedSettlementDate { get; set; }

    [ProtoMember(13)]
    public DateTime ForecastedSettlementDate { get; set; }

    [ProtoMember(14)]
    public DateTime ScheduledExecutionTimestamp { get; set; }

    [ProtoMember(15)]
    public string AppId { get; set; }

    [ProtoMember(16)]
    public string AppQueue { get; set; }

    [ProtoMember(17)]
    public string PaymentMethod { get; set; }

    public ScheduleInstructionStep3V2()
    {
    }

    public ScheduleInstructionStep3V2(string correlationId, Guid requestId, Guid tenantId, string customerRef, Guid financialInstrumentId, string workflowId, string orderRef, string instructionRef, string country, string currency, decimal amount, DateTime requestedSettlementDate, DateTime forecastedSettlementDate, DateTime scheduledExecutionTimestamp, string appId, string appQueue, string paymentMethod)
    {
        CorrelationId = correlationId;
        RequestId = requestId;
        TenantId = tenantId;
        CustomerRef = customerRef;
        FinancialInstrumentId = financialInstrumentId;
        WorkflowId = workflowId;
        OrderRef = orderRef;
        InstructionRef = instructionRef;
        Country = country;
        Currency = currency;
        Amount = amount;
        RequestedSettlementDate = requestedSettlementDate;
        ForecastedSettlementDate = forecastedSettlementDate;
        ScheduledExecutionTimestamp = scheduledExecutionTimestamp;
        AppId = appId;
        AppQueue = appQueue;
        PaymentMethod = paymentMethod;
    }
}