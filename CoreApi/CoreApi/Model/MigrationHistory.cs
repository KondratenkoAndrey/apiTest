﻿using System;
using System.Collections.Generic;

namespace CoreApi.Model
{
    public partial class MigrationHistory
    {
        public string MigrationId { get; set; }
        public string ContextKey { get; set; }
        public byte[] Model { get; set; }
        public string ProductVersion { get; set; }
    }
}
